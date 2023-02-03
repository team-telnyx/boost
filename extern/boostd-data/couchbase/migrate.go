package couchbase

import (
	"context"
	"fmt"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
	"go.uber.org/zap"
	"strings"
)

func (s *Store) Migrate(ctx context.Context, logger *zap.SugaredLogger, rowCount int) error {
	err := s.migratePieceCidToMetadata(ctx, logger, rowCount)
	if err != nil {
		return fmt.Errorf("migrating piece cid to metadata bucket: %w", err)
	}

	return nil
}

type pcidRow struct {
	RowID     string
	PieceMeta CouchbaseMetadata
}

func (s *Store) migratePieceCidToMetadata(ctx context.Context, logger *zap.SugaredLogger, rowCount int) error {
	tableName := "`" + pieceCidToMetadataBucket + "`"
	qry := "SELECT META(PieceMeta).id as RowID, * FROM " + tableName + " as PieceMeta"
	res, err := s.db.query(ctx, qry)
	if err != nil {
		return fmt.Errorf("getting keys from %s: %w", tableName, err)
	}

	var rows []pcidRow
	for i := 0; i < rowCount && res.Next(); i++ {
		var r pcidRow
		err := res.Row(&r)
		if err != nil {
			return fmt.Errorf("reading row from %s: %w", pieceCidToMetadataBucket, err)
		}

		rows = append(rows, r)
	}

	err = res.Close()
	if err != nil {
		return fmt.Errorf("closing stream from %s: %w", pieceCidToMetadataBucket, err)
	}

	logger.Infof("migrating %d rows", len(rows))

	for i, r := range rows {
		rowId := strings.Replace(r.RowID, "u:", "", 1)
		pieceCid, err := cid.Parse(rowId)
		if err != nil {
			logger.Errorf("parsing row id %s as piece cid: %s", r.RowID, err)
			continue
		}

		err = s.migrate(ctx, pieceCid, r)
		if err != nil {
			logger.Errorf("migrating piece %s: %s", r.RowID, err)
			continue
		}

		r.PieceMeta.Version = pieceMetadataVersion
		err = s.db.mutatePieceMetadata(ctx, pieceCid, "migrate-piece-metadata", func(metadata CouchbaseMetadata) *CouchbaseMetadata {
			return &r.PieceMeta
		})
		if err != nil {
			logger.Errorf("mutating metadata for piece %s: %s", pieceCid, err)
			continue
		}

		logger.Infow("migrated piece", "piece cid", pieceCid, "row", i, "total", len(rows))
	}

	logger.Infof("migrated %d rows", len(rows))

	return nil
}

func (s *Store) migrate(ctx context.Context, pieceCid cid.Cid, r pcidRow) error {
	recs, err := s.db.allRecordsWithKeyFn(ctx, pieceCid, r.PieceMeta.BlockCount, func(k string) string {
		return toCouchKey("u:" + k)
	}, multihash.FromHexString)
	if err != nil {
		return fmt.Errorf("getting recs: %w", err)
	}

	err = s.AddIndex(ctx, pieceCid, recs, r.PieceMeta.CompleteIndex)
	if err != nil {
		return fmt.Errorf("adding index: %w", err)
	}

	return nil
}

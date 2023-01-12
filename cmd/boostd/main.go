package main

import (
	"os"

	"github.com/filecoin-project/boost/cmd"
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/boost/build"
	"github.com/filecoin-project/boost/cli/unchained"
	cliutil "github.com/filecoin-project/boost/cli/util"
)

var log = logging.Logger("boostd")

const (
	FlagBoostRepo = "boost-repo"
	FlagUnchained = "boost-standalone"
)

func main() {
	app := &cli.App{
		Name:                 "boostd",
		Usage:                "Markets V2 module for Filecoin",
		EnableBashCompletion: true,
		Version:              build.UserVersion(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    FlagBoostRepo,
				EnvVars: []string{"BOOST_PATH"},
				Usage:   "boost repo path",
				Value:   "~/.boost",
			},
			&cli.BoolFlag{
				Name:  FlagUnchained,
				Usage: "run boost in retrieval mode without a local lotus daemon",
				Value: false,
			},
			cmd.FlagJson,
			cliutil.FlagVeryVerbose,
		},
		Commands: []*cli.Command{
			authCmd,
			runCmd,
			initCmd,
			migrateMonolithCmd,
			migrateMarketsCmd,
			backupCmd,
			restoreCmd,
			dummydealCmd,
			dataTransfersCmd,
			retrievalDealsCmd,
			indexProvCmd,
			importDataCmd,
			logCmd,
			dagstoreCmd,
			piecesCmd,
			netCmd,
		},
	}
	app.Setup()

	if err := app.Run(os.Args); err != nil {
		os.Stderr.WriteString("Error: " + err.Error() + "\n")
	}
}

func before(cctx *cli.Context) error {
	_ = logging.SetLogLevel("boostd", "INFO")
	_ = logging.SetLogLevel("db", "INFO")
	_ = logging.SetLogLevel("boost-prop", "INFO")
	_ = logging.SetLogLevel("modules", "INFO")
	_ = logging.SetLogLevel("cfg", "INFO")
	_ = logging.SetLogLevel("boost-storage-deal", "INFO")

	if cctx.Bool(FlagUnchained) {
		cctx.App.Metadata["testnode-full"] = unchained.MakeProxy(cctx)
		cctx.App.After = func(ctx *cli.Context) error {
			v1c := cctx.App.Metadata["testnode-full"].(*unchained.Node)
			v1c.Close()
			return nil
		}
	}

	if cliutil.IsVeryVerbose {
		_ = logging.SetLogLevel("boostd", "DEBUG")
		_ = logging.SetLogLevel("provider", "DEBUG")
		_ = logging.SetLogLevel("boost-net", "DEBUG")
		_ = logging.SetLogLevel("boost-provider", "DEBUG")
		_ = logging.SetLogLevel("storagemanager", "DEBUG")
		_ = logging.SetLogLevel("index-provider-wrapper", "DEBUG")
		_ = logging.SetLogLevel("boost-migrator", "DEBUG")
		_ = logging.SetLogLevel("dagstore", "DEBUG")
		_ = logging.SetLogLevel("migrator", "DEBUG")
	}

	return nil
}

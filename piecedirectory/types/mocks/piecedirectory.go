// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/filecoin-project/boost/piecedirectory/types (interfaces: SectionReader,PieceReader,Store)

// Package mock_piecedirectory is a generated GoMock package.
package mock_piecedirectory

import (
	context "context"
	reflect "reflect"
	time "time"

	types "github.com/filecoin-project/boost/piecedirectory/types"
	model "github.com/filecoin-project/boostd-data/model"
	abi "github.com/filecoin-project/go-state-types/abi"
	gomock "github.com/golang/mock/gomock"
	cid "github.com/ipfs/go-cid"
	index "github.com/ipld/go-car/v2/index"
	multihash "github.com/multiformats/go-multihash"
)

// MockSectionReader is a mock of SectionReader interface.
type MockSectionReader struct {
	ctrl     *gomock.Controller
	recorder *MockSectionReaderMockRecorder
}

// MockSectionReaderMockRecorder is the mock recorder for MockSectionReader.
type MockSectionReaderMockRecorder struct {
	mock *MockSectionReader
}

// NewMockSectionReader creates a new mock instance.
func NewMockSectionReader(ctrl *gomock.Controller) *MockSectionReader {
	mock := &MockSectionReader{ctrl: ctrl}
	mock.recorder = &MockSectionReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSectionReader) EXPECT() *MockSectionReaderMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockSectionReader) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSectionReaderMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSectionReader)(nil).Close))
}

// Read mocks base method.
func (m *MockSectionReader) Read(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockSectionReaderMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockSectionReader)(nil).Read), arg0)
}

// ReadAt mocks base method.
func (m *MockSectionReader) ReadAt(arg0 []byte, arg1 int64) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAt", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAt indicates an expected call of ReadAt.
func (mr *MockSectionReaderMockRecorder) ReadAt(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAt", reflect.TypeOf((*MockSectionReader)(nil).ReadAt), arg0, arg1)
}

// Seek mocks base method.
func (m *MockSectionReader) Seek(arg0 int64, arg1 int) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seek", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Seek indicates an expected call of Seek.
func (mr *MockSectionReaderMockRecorder) Seek(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seek", reflect.TypeOf((*MockSectionReader)(nil).Seek), arg0, arg1)
}

// MockPieceReader is a mock of PieceReader interface.
type MockPieceReader struct {
	ctrl     *gomock.Controller
	recorder *MockPieceReaderMockRecorder
}

// MockPieceReaderMockRecorder is the mock recorder for MockPieceReader.
type MockPieceReaderMockRecorder struct {
	mock *MockPieceReader
}

// NewMockPieceReader creates a new mock instance.
func NewMockPieceReader(ctrl *gomock.Controller) *MockPieceReader {
	mock := &MockPieceReader{ctrl: ctrl}
	mock.recorder = &MockPieceReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPieceReader) EXPECT() *MockPieceReaderMockRecorder {
	return m.recorder
}

// GetReader mocks base method.
func (m *MockPieceReader) GetReader(arg0 context.Context, arg1 abi.SectorNumber, arg2, arg3 abi.PaddedPieceSize) (types.SectionReader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReader", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(types.SectionReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReader indicates an expected call of GetReader.
func (mr *MockPieceReaderMockRecorder) GetReader(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReader", reflect.TypeOf((*MockPieceReader)(nil).GetReader), arg0, arg1, arg2, arg3)
}

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddDealForPiece mocks base method.
func (m *MockStore) AddDealForPiece(arg0 context.Context, arg1 cid.Cid, arg2 model.DealInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDealForPiece", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDealForPiece indicates an expected call of AddDealForPiece.
func (mr *MockStoreMockRecorder) AddDealForPiece(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDealForPiece", reflect.TypeOf((*MockStore)(nil).AddDealForPiece), arg0, arg1, arg2)
}

// AddIndex mocks base method.
func (m *MockStore) AddIndex(arg0 context.Context, arg1 cid.Cid, arg2 []model.Record, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddIndex", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddIndex indicates an expected call of AddIndex.
func (mr *MockStoreMockRecorder) AddIndex(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIndex", reflect.TypeOf((*MockStore)(nil).AddIndex), arg0, arg1, arg2, arg3)
}

// FlagPiece mocks base method.
func (m *MockStore) FlagPiece(arg0 context.Context, arg1 cid.Cid) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlagPiece", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// FlagPiece indicates an expected call of FlagPiece.
func (mr *MockStoreMockRecorder) FlagPiece(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlagPiece", reflect.TypeOf((*MockStore)(nil).FlagPiece), arg0, arg1)
}

// FlaggedPiecesCount mocks base method.
func (m *MockStore) FlaggedPiecesCount(arg0 context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlaggedPiecesCount", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FlaggedPiecesCount indicates an expected call of FlaggedPiecesCount.
func (mr *MockStoreMockRecorder) FlaggedPiecesCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlaggedPiecesCount", reflect.TypeOf((*MockStore)(nil).FlaggedPiecesCount), arg0)
}

// FlaggedPiecesList mocks base method.
func (m *MockStore) FlaggedPiecesList(arg0 context.Context, arg1 *time.Time, arg2, arg3 int) ([]model.FlaggedPiece, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlaggedPiecesList", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]model.FlaggedPiece)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FlaggedPiecesList indicates an expected call of FlaggedPiecesList.
func (mr *MockStoreMockRecorder) FlaggedPiecesList(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlaggedPiecesList", reflect.TypeOf((*MockStore)(nil).FlaggedPiecesList), arg0, arg1, arg2, arg3)
}

// GetIndex mocks base method.
func (m *MockStore) GetIndex(arg0 context.Context, arg1 cid.Cid) (index.Index, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIndex", arg0, arg1)
	ret0, _ := ret[0].(index.Index)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIndex indicates an expected call of GetIndex.
func (mr *MockStoreMockRecorder) GetIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIndex", reflect.TypeOf((*MockStore)(nil).GetIndex), arg0, arg1)
}

// GetOffsetSize mocks base method.
func (m *MockStore) GetOffsetSize(arg0 context.Context, arg1 cid.Cid, arg2 multihash.Multihash) (*model.OffsetSize, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOffsetSize", arg0, arg1, arg2)
	ret0, _ := ret[0].(*model.OffsetSize)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOffsetSize indicates an expected call of GetOffsetSize.
func (mr *MockStoreMockRecorder) GetOffsetSize(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOffsetSize", reflect.TypeOf((*MockStore)(nil).GetOffsetSize), arg0, arg1, arg2)
}

// GetPieceDeals mocks base method.
func (m *MockStore) GetPieceDeals(arg0 context.Context, arg1 cid.Cid) ([]model.DealInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPieceDeals", arg0, arg1)
	ret0, _ := ret[0].([]model.DealInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPieceDeals indicates an expected call of GetPieceDeals.
func (mr *MockStoreMockRecorder) GetPieceDeals(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPieceDeals", reflect.TypeOf((*MockStore)(nil).GetPieceDeals), arg0, arg1)
}

// GetPieceMetadata mocks base method.
func (m *MockStore) GetPieceMetadata(arg0 context.Context, arg1 cid.Cid) (model.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPieceMetadata", arg0, arg1)
	ret0, _ := ret[0].(model.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPieceMetadata indicates an expected call of GetPieceMetadata.
func (mr *MockStoreMockRecorder) GetPieceMetadata(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPieceMetadata", reflect.TypeOf((*MockStore)(nil).GetPieceMetadata), arg0, arg1)
}

// IsCompleteIndex mocks base method.
func (m *MockStore) IsCompleteIndex(arg0 context.Context, arg1 cid.Cid) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCompleteIndex", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsCompleteIndex indicates an expected call of IsCompleteIndex.
func (mr *MockStoreMockRecorder) IsCompleteIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCompleteIndex", reflect.TypeOf((*MockStore)(nil).IsCompleteIndex), arg0, arg1)
}

// IsIndexed mocks base method.
func (m *MockStore) IsIndexed(arg0 context.Context, arg1 cid.Cid) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsIndexed", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsIndexed indicates an expected call of IsIndexed.
func (mr *MockStoreMockRecorder) IsIndexed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsIndexed", reflect.TypeOf((*MockStore)(nil).IsIndexed), arg0, arg1)
}

// ListPieces mocks base method.
func (m *MockStore) ListPieces(arg0 context.Context) ([]cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPieces", arg0)
	ret0, _ := ret[0].([]cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPieces indicates an expected call of ListPieces.
func (mr *MockStoreMockRecorder) ListPieces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPieces", reflect.TypeOf((*MockStore)(nil).ListPieces), arg0)
}

// NextPiecesToCheck mocks base method.
func (m *MockStore) NextPiecesToCheck(arg0 context.Context) ([]cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextPiecesToCheck", arg0)
	ret0, _ := ret[0].([]cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NextPiecesToCheck indicates an expected call of NextPiecesToCheck.
func (mr *MockStoreMockRecorder) NextPiecesToCheck(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextPiecesToCheck", reflect.TypeOf((*MockStore)(nil).NextPiecesToCheck), arg0)
}

// PiecesContainingMultihash mocks base method.
func (m *MockStore) PiecesContainingMultihash(arg0 context.Context, arg1 multihash.Multihash) ([]cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PiecesContainingMultihash", arg0, arg1)
	ret0, _ := ret[0].([]cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PiecesContainingMultihash indicates an expected call of PiecesContainingMultihash.
func (mr *MockStoreMockRecorder) PiecesContainingMultihash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PiecesContainingMultihash", reflect.TypeOf((*MockStore)(nil).PiecesContainingMultihash), arg0, arg1)
}

// RemoveDealForPiece mocks base method.
func (m *MockStore) RemoveDealForPiece(arg0 context.Context, arg1 cid.Cid, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveDealForPiece", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveDealForPiece indicates an expected call of RemoveDealForPiece.
func (mr *MockStoreMockRecorder) RemoveDealForPiece(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDealForPiece", reflect.TypeOf((*MockStore)(nil).RemoveDealForPiece), arg0, arg1, arg2)
}

// RemoveIndexes mocks base method.
func (m *MockStore) RemoveIndexes(arg0 context.Context, arg1 cid.Cid) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveIndexes", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveIndexes indicates an expected call of RemoveIndexes.
func (mr *MockStoreMockRecorder) RemoveIndexes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveIndexes", reflect.TypeOf((*MockStore)(nil).RemoveIndexes), arg0, arg1)
}

// RemovePieceMetadata mocks base method.
func (m *MockStore) RemovePieceMetadata(arg0 context.Context, arg1 cid.Cid) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePieceMetadata", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePieceMetadata indicates an expected call of RemovePieceMetadata.
func (mr *MockStoreMockRecorder) RemovePieceMetadata(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePieceMetadata", reflect.TypeOf((*MockStore)(nil).RemovePieceMetadata), arg0, arg1)
}

// UnflagPiece mocks base method.
func (m *MockStore) UnflagPiece(arg0 context.Context, arg1 cid.Cid) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnflagPiece", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnflagPiece indicates an expected call of UnflagPiece.
func (mr *MockStoreMockRecorder) UnflagPiece(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnflagPiece", reflect.TypeOf((*MockStore)(nil).UnflagPiece), arg0, arg1)
}

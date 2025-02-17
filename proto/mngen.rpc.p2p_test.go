// Code generated. DO NOT EDIT.
package proto

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testP2PStatusHandler struct{}

func (th *testP2PStatusHandler) HandleP2PStatus(context.Context, *StatusRequest) (*StatusResponse, error) {
	return &StatusResponse{
		SyncToBlockHeight:  1,
		MaxBlockHeightSeen: 2,
	}, nil
}

func TestP2PStatus(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PStatusHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PStatus(h)

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	res, err := srvr.Status(context.Background(), &StatusRequest{})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1, res.SyncToBlockHeight, "the sync to block height must match")
	assert.Equal(t, 2, res.MaxBlockHeightSeen, "the max block height seen must match")
}

func TestDoubleregistrationP2PStatus(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PStatusHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PStatus(h)

	fn := func() {
		d.RegisterP2PStatus(h)
	}
	assert.Panics(t, fn, "double registration must panic")
}

func TestP2PStatusCancel(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	errChan := make(chan error)
	defer close(errChan)
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	fn := func() {
		_, err := srvr.Status(cancelCtx, &StatusRequest{})
		errChan <- err
	}
	go fn()
	cancelFunc()
	cancelErr := <-errChan
	assert.EqualError(t, cancelErr, "context canceled", "the error returned must be a context canceled error")
}

type testP2PGetBlockHeadersHandler struct{}

func (th *testP2PGetBlockHeadersHandler) HandleP2PGetBlockHeaders(context.Context, *GetBlockHeadersRequest) (*GetBlockHeadersResponse, error) {
	return &GetBlockHeadersResponse{}, nil
}

func TestP2PGetBlockHeaders(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGetBlockHeadersHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGetBlockHeaders(h)

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	_, err := srvr.GetBlockHeaders(context.Background(), &GetBlockHeadersRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestDoubleregistrationP2PGetBlockHeaders(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGetBlockHeadersHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGetBlockHeaders(h)

	fn := func() {
		d.RegisterP2PGetBlockHeaders(h)
	}
	assert.Panics(t, fn, "double registration must panic")
}

func TestP2PGetBlockHeadersCancel(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	errChan := make(chan error)
	defer close(errChan)
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	fn := func() {
		_, err := srvr.GetBlockHeaders(cancelCtx, &GetBlockHeadersRequest{})
		errChan <- err
	}
	go fn()
	cancelFunc()
	cancelErr := <-errChan
	assert.EqualError(t, cancelErr, "context canceled", "the error returned must be a context canceled error")
}

type testP2PGetMinedTxsHandler struct{}

func (th *testP2PGetMinedTxsHandler) HandleP2PGetMinedTxs(context.Context, *GetMinedTxsRequest) (*GetMinedTxsResponse, error) {
	return &GetMinedTxsResponse{}, nil
}

func TestP2PGetMinedTxs(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGetMinedTxsHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGetMinedTxs(h)

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	_, err := srvr.GetMinedTxs(context.Background(), &GetMinedTxsRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestDoubleregistrationP2PGetMinedTxs(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGetMinedTxsHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGetMinedTxs(h)

	fn := func() {
		d.RegisterP2PGetMinedTxs(h)
	}
	assert.Panics(t, fn, "double registration must panic")
}

func TestP2PGetMinedTxsCancel(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	errChan := make(chan error)
	defer close(errChan)
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	fn := func() {
		_, err := srvr.GetMinedTxs(cancelCtx, &GetMinedTxsRequest{})
		errChan <- err
	}
	go fn()
	cancelFunc()
	cancelErr := <-errChan
	assert.EqualError(t, cancelErr, "context canceled", "the error returned must be a context canceled error")
}

type testP2PGetPendingTxsHandler struct{}

func (th *testP2PGetPendingTxsHandler) HandleP2PGetPendingTxs(context.Context, *GetPendingTxsRequest) (*GetPendingTxsResponse, error) {
	return &GetPendingTxsResponse{}, nil
}

func TestP2PGetPendingTxs(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGetPendingTxsHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGetPendingTxs(h)

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	_, err := srvr.GetPendingTxs(context.Background(), &GetPendingTxsRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestDoubleregistrationP2PGetPendingTxs(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGetPendingTxsHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGetPendingTxs(h)

	fn := func() {
		d.RegisterP2PGetPendingTxs(h)
	}
	assert.Panics(t, fn, "double registration must panic")
}

func TestP2PGetPendingTxsCancel(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	errChan := make(chan error)
	defer close(errChan)
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	fn := func() {
		_, err := srvr.GetPendingTxs(cancelCtx, &GetPendingTxsRequest{})
		errChan <- err
	}
	go fn()
	cancelFunc()
	cancelErr := <-errChan
	assert.EqualError(t, cancelErr, "context canceled", "the error returned must be a context canceled error")
}

type testP2PGetSnapShotNodeHandler struct{}

func (th *testP2PGetSnapShotNodeHandler) HandleP2PGetSnapShotNode(context.Context, *GetSnapShotNodeRequest) (*GetSnapShotNodeResponse, error) {
	return &GetSnapShotNodeResponse{}, nil
}

func TestP2PGetSnapShotNode(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGetSnapShotNodeHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGetSnapShotNode(h)

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	_, err := srvr.GetSnapShotNode(context.Background(), &GetSnapShotNodeRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestDoubleregistrationP2PGetSnapShotNode(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGetSnapShotNodeHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGetSnapShotNode(h)

	fn := func() {
		d.RegisterP2PGetSnapShotNode(h)
	}
	assert.Panics(t, fn, "double registration must panic")
}

func TestP2PGetSnapShotNodeCancel(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	errChan := make(chan error)
	defer close(errChan)
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	fn := func() {
		_, err := srvr.GetSnapShotNode(cancelCtx, &GetSnapShotNodeRequest{})
		errChan <- err
	}
	go fn()
	cancelFunc()
	cancelErr := <-errChan
	assert.EqualError(t, cancelErr, "context canceled", "the error returned must be a context canceled error")
}

type testP2PGetSnapShotStateDataHandler struct{}

func (th *testP2PGetSnapShotStateDataHandler) HandleP2PGetSnapShotStateData(context.Context, *GetSnapShotStateDataRequest) (*GetSnapShotStateDataResponse, error) {
	return &GetSnapShotStateDataResponse{}, nil
}

func TestP2PGetSnapShotStateData(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGetSnapShotStateDataHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGetSnapShotStateData(h)

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	_, err := srvr.GetSnapShotStateData(context.Background(), &GetSnapShotStateDataRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestDoubleregistrationP2PGetSnapShotStateData(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGetSnapShotStateDataHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGetSnapShotStateData(h)

	fn := func() {
		d.RegisterP2PGetSnapShotStateData(h)
	}
	assert.Panics(t, fn, "double registration must panic")
}

func TestP2PGetSnapShotStateDataCancel(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	errChan := make(chan error)
	defer close(errChan)
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	fn := func() {
		_, err := srvr.GetSnapShotStateData(cancelCtx, &GetSnapShotStateDataRequest{})
		errChan <- err
	}
	go fn()
	cancelFunc()
	cancelErr := <-errChan
	assert.EqualError(t, cancelErr, "context canceled", "the error returned must be a context canceled error")
}

type testP2PGetSnapShotHdrNodeHandler struct{}

func (th *testP2PGetSnapShotHdrNodeHandler) HandleP2PGetSnapShotHdrNode(context.Context, *GetSnapShotHdrNodeRequest) (*GetSnapShotHdrNodeResponse, error) {
	return &GetSnapShotHdrNodeResponse{}, nil
}

func TestP2PGetSnapShotHdrNode(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGetSnapShotHdrNodeHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGetSnapShotHdrNode(h)

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	_, err := srvr.GetSnapShotHdrNode(context.Background(), &GetSnapShotHdrNodeRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestDoubleregistrationP2PGetSnapShotHdrNode(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGetSnapShotHdrNodeHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGetSnapShotHdrNode(h)

	fn := func() {
		d.RegisterP2PGetSnapShotHdrNode(h)
	}
	assert.Panics(t, fn, "double registration must panic")
}

func TestP2PGetSnapShotHdrNodeCancel(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	errChan := make(chan error)
	defer close(errChan)
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	fn := func() {
		_, err := srvr.GetSnapShotHdrNode(cancelCtx, &GetSnapShotHdrNodeRequest{})
		errChan <- err
	}
	go fn()
	cancelFunc()
	cancelErr := <-errChan
	assert.EqualError(t, cancelErr, "context canceled", "the error returned must be a context canceled error")
}

type testP2PGossipTransactionHandler struct{}

func (th *testP2PGossipTransactionHandler) HandleP2PGossipTransaction(context.Context, *GossipTransactionMessage) (*GossipTransactionAck, error) {
	return &GossipTransactionAck{}, nil
}

func TestP2PGossipTransaction(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGossipTransactionHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGossipTransaction(h)

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	_, err := srvr.GossipTransaction(context.Background(), &GossipTransactionMessage{})
	if err != nil {
		t.Error(err)
	}
}

func TestDoubleregistrationP2PGossipTransaction(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGossipTransactionHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGossipTransaction(h)

	fn := func() {
		d.RegisterP2PGossipTransaction(h)
	}
	assert.Panics(t, fn, "double registration must panic")
}

func TestP2PGossipTransactionCancel(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	errChan := make(chan error)
	defer close(errChan)
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	fn := func() {
		_, err := srvr.GossipTransaction(cancelCtx, &GossipTransactionMessage{})
		errChan <- err
	}
	go fn()
	cancelFunc()
	cancelErr := <-errChan
	assert.EqualError(t, cancelErr, "context canceled", "the error returned must be a context canceled error")
}

type testP2PGossipProposalHandler struct{}

func (th *testP2PGossipProposalHandler) HandleP2PGossipProposal(context.Context, *GossipProposalMessage) (*GossipProposalAck, error) {
	return &GossipProposalAck{}, nil
}

func TestP2PGossipProposal(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGossipProposalHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGossipProposal(h)

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	_, err := srvr.GossipProposal(context.Background(), &GossipProposalMessage{})
	if err != nil {
		t.Error(err)
	}
}

func TestDoubleregistrationP2PGossipProposal(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGossipProposalHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGossipProposal(h)

	fn := func() {
		d.RegisterP2PGossipProposal(h)
	}
	assert.Panics(t, fn, "double registration must panic")
}

func TestP2PGossipProposalCancel(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	errChan := make(chan error)
	defer close(errChan)
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	fn := func() {
		_, err := srvr.GossipProposal(cancelCtx, &GossipProposalMessage{})
		errChan <- err
	}
	go fn()
	cancelFunc()
	cancelErr := <-errChan
	assert.EqualError(t, cancelErr, "context canceled", "the error returned must be a context canceled error")
}

type testP2PGossipPreVoteHandler struct{}

func (th *testP2PGossipPreVoteHandler) HandleP2PGossipPreVote(context.Context, *GossipPreVoteMessage) (*GossipPreVoteAck, error) {
	return &GossipPreVoteAck{}, nil
}

func TestP2PGossipPreVote(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGossipPreVoteHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGossipPreVote(h)

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	_, err := srvr.GossipPreVote(context.Background(), &GossipPreVoteMessage{})
	if err != nil {
		t.Error(err)
	}
}

func TestDoubleregistrationP2PGossipPreVote(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGossipPreVoteHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGossipPreVote(h)

	fn := func() {
		d.RegisterP2PGossipPreVote(h)
	}
	assert.Panics(t, fn, "double registration must panic")
}

func TestP2PGossipPreVoteCancel(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	errChan := make(chan error)
	defer close(errChan)
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	fn := func() {
		_, err := srvr.GossipPreVote(cancelCtx, &GossipPreVoteMessage{})
		errChan <- err
	}
	go fn()
	cancelFunc()
	cancelErr := <-errChan
	assert.EqualError(t, cancelErr, "context canceled", "the error returned must be a context canceled error")
}

type testP2PGossipPreVoteNilHandler struct{}

func (th *testP2PGossipPreVoteNilHandler) HandleP2PGossipPreVoteNil(context.Context, *GossipPreVoteNilMessage) (*GossipPreVoteNilAck, error) {
	return &GossipPreVoteNilAck{}, nil
}

func TestP2PGossipPreVoteNil(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGossipPreVoteNilHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGossipPreVoteNil(h)

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	_, err := srvr.GossipPreVoteNil(context.Background(), &GossipPreVoteNilMessage{})
	if err != nil {
		t.Error(err)
	}
}

func TestDoubleregistrationP2PGossipPreVoteNil(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGossipPreVoteNilHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGossipPreVoteNil(h)

	fn := func() {
		d.RegisterP2PGossipPreVoteNil(h)
	}
	assert.Panics(t, fn, "double registration must panic")
}

func TestP2PGossipPreVoteNilCancel(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	errChan := make(chan error)
	defer close(errChan)
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	fn := func() {
		_, err := srvr.GossipPreVoteNil(cancelCtx, &GossipPreVoteNilMessage{})
		errChan <- err
	}
	go fn()
	cancelFunc()
	cancelErr := <-errChan
	assert.EqualError(t, cancelErr, "context canceled", "the error returned must be a context canceled error")
}

type testP2PGossipPreCommitHandler struct{}

func (th *testP2PGossipPreCommitHandler) HandleP2PGossipPreCommit(context.Context, *GossipPreCommitMessage) (*GossipPreCommitAck, error) {
	return &GossipPreCommitAck{}, nil
}

func TestP2PGossipPreCommit(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGossipPreCommitHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGossipPreCommit(h)

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	_, err := srvr.GossipPreCommit(context.Background(), &GossipPreCommitMessage{})
	if err != nil {
		t.Error(err)
	}
}

func TestDoubleregistrationP2PGossipPreCommit(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGossipPreCommitHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGossipPreCommit(h)

	fn := func() {
		d.RegisterP2PGossipPreCommit(h)
	}
	assert.Panics(t, fn, "double registration must panic")
}

func TestP2PGossipPreCommitCancel(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	errChan := make(chan error)
	defer close(errChan)
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	fn := func() {
		_, err := srvr.GossipPreCommit(cancelCtx, &GossipPreCommitMessage{})
		errChan <- err
	}
	go fn()
	cancelFunc()
	cancelErr := <-errChan
	assert.EqualError(t, cancelErr, "context canceled", "the error returned must be a context canceled error")
}

type testP2PGossipPreCommitNilHandler struct{}

func (th *testP2PGossipPreCommitNilHandler) HandleP2PGossipPreCommitNil(context.Context, *GossipPreCommitNilMessage) (*GossipPreCommitNilAck, error) {
	return &GossipPreCommitNilAck{}, nil
}

func TestP2PGossipPreCommitNil(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGossipPreCommitNilHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGossipPreCommitNil(h)

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	_, err := srvr.GossipPreCommitNil(context.Background(), &GossipPreCommitNilMessage{})
	if err != nil {
		t.Error(err)
	}
}

func TestDoubleregistrationP2PGossipPreCommitNil(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGossipPreCommitNilHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGossipPreCommitNil(h)

	fn := func() {
		d.RegisterP2PGossipPreCommitNil(h)
	}
	assert.Panics(t, fn, "double registration must panic")
}

func TestP2PGossipPreCommitNilCancel(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	errChan := make(chan error)
	defer close(errChan)
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	fn := func() {
		_, err := srvr.GossipPreCommitNil(cancelCtx, &GossipPreCommitNilMessage{})
		errChan <- err
	}
	go fn()
	cancelFunc()
	cancelErr := <-errChan
	assert.EqualError(t, cancelErr, "context canceled", "the error returned must be a context canceled error")
}

type testP2PGossipNextRoundHandler struct{}

func (th *testP2PGossipNextRoundHandler) HandleP2PGossipNextRound(context.Context, *GossipNextRoundMessage) (*GossipNextRoundAck, error) {
	return &GossipNextRoundAck{}, nil
}

func TestP2PGossipNextRound(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGossipNextRoundHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGossipNextRound(h)

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	_, err := srvr.GossipNextRound(context.Background(), &GossipNextRoundMessage{})
	if err != nil {
		t.Error(err)
	}
}

func TestDoubleregistrationP2PGossipNextRound(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGossipNextRoundHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGossipNextRound(h)

	fn := func() {
		d.RegisterP2PGossipNextRound(h)
	}
	assert.Panics(t, fn, "double registration must panic")
}

func TestP2PGossipNextRoundCancel(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	errChan := make(chan error)
	defer close(errChan)
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	fn := func() {
		_, err := srvr.GossipNextRound(cancelCtx, &GossipNextRoundMessage{})
		errChan <- err
	}
	go fn()
	cancelFunc()
	cancelErr := <-errChan
	assert.EqualError(t, cancelErr, "context canceled", "the error returned must be a context canceled error")
}

type testP2PGossipNextHeightHandler struct{}

func (th *testP2PGossipNextHeightHandler) HandleP2PGossipNextHeight(context.Context, *GossipNextHeightMessage) (*GossipNextHeightAck, error) {
	return &GossipNextHeightAck{}, nil
}

func TestP2PGossipNextHeight(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGossipNextHeightHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGossipNextHeight(h)

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	_, err := srvr.GossipNextHeight(context.Background(), &GossipNextHeightMessage{})
	if err != nil {
		t.Error(err)
	}
}

func TestDoubleregistrationP2PGossipNextHeight(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGossipNextHeightHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGossipNextHeight(h)

	fn := func() {
		d.RegisterP2PGossipNextHeight(h)
	}
	assert.Panics(t, fn, "double registration must panic")
}

func TestP2PGossipNextHeightCancel(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	errChan := make(chan error)
	defer close(errChan)
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	fn := func() {
		_, err := srvr.GossipNextHeight(cancelCtx, &GossipNextHeightMessage{})
		errChan <- err
	}
	go fn()
	cancelFunc()
	cancelErr := <-errChan
	assert.EqualError(t, cancelErr, "context canceled", "the error returned must be a context canceled error")
}

type testP2PGossipBlockHeaderHandler struct{}

func (th *testP2PGossipBlockHeaderHandler) HandleP2PGossipBlockHeader(context.Context, *GossipBlockHeaderMessage) (*GossipBlockHeaderAck, error) {
	return &GossipBlockHeaderAck{}, nil
}

func TestP2PGossipBlockHeader(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGossipBlockHeaderHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGossipBlockHeader(h)

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	_, err := srvr.GossipBlockHeader(context.Background(), &GossipBlockHeaderMessage{})
	if err != nil {
		t.Error(err)
	}
}

func TestDoubleregistrationP2PGossipBlockHeader(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGossipBlockHeaderHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGossipBlockHeader(h)

	fn := func() {
		d.RegisterP2PGossipBlockHeader(h)
	}
	assert.Panics(t, fn, "double registration must panic")
}

func TestP2PGossipBlockHeaderCancel(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	errChan := make(chan error)
	defer close(errChan)
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	fn := func() {
		_, err := srvr.GossipBlockHeader(cancelCtx, &GossipBlockHeaderMessage{})
		errChan <- err
	}
	go fn()
	cancelFunc()
	cancelErr := <-errChan
	assert.EqualError(t, cancelErr, "context canceled", "the error returned must be a context canceled error")
}

type testP2PGetPeersHandler struct{}

func (th *testP2PGetPeersHandler) HandleP2PGetPeers(context.Context, *GetPeersRequest) (*GetPeersResponse, error) {
	return &GetPeersResponse{}, nil
}

func TestP2PGetPeers(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGetPeersHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGetPeers(h)

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	_, err := srvr.GetPeers(context.Background(), &GetPeersRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestDoubleregistrationP2PGetPeers(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testP2PGetPeersHandler{}

	// Register the handler with the dispatch class
	d.RegisterP2PGetPeers(h)

	fn := func() {
		d.RegisterP2PGetPeers(h)
	}
	assert.Panics(t, fn, "double registration must panic")
}

func TestP2PGetPeersCancel(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Create the server and pass in the dispatch class
	srvr := GeneratedP2PServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	errChan := make(chan error)
	defer close(errChan)
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	fn := func() {
		_, err := srvr.GetPeers(cancelCtx, &GetPeersRequest{})
		errChan <- err
	}
	go fn()
	cancelFunc()
	cancelErr := <-errChan
	assert.EqualError(t, cancelErr, "context canceled", "the error returned must be a context canceled error")
}

type testDiscoveryGetPeersHandler struct{}

func (th *testDiscoveryGetPeersHandler) HandleDiscoveryGetPeers(context.Context, *GetPeersRequest) (*GetPeersResponse, error) {
	return &GetPeersResponse{}, nil
}

func TestDiscoveryGetPeers(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testDiscoveryGetPeersHandler{}

	// Register the handler with the dispatch class
	d.RegisterDiscoveryGetPeers(h)

	// Create the server and pass in the dispatch class
	srvr := GeneratedDiscoveryServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	_, err := srvr.GetPeers(context.Background(), &GetPeersRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestDoubleregistrationDiscoveryGetPeers(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Setup the handler for the TestService
	h := &testDiscoveryGetPeersHandler{}

	// Register the handler with the dispatch class
	d.RegisterDiscoveryGetPeers(h)

	fn := func() {
		d.RegisterDiscoveryGetPeers(h)
	}
	assert.Panics(t, fn, "double registration must panic")
}

func TestDiscoveryGetPeersCancel(t *testing.T) {
	// Setup the dispatch handler
	d := NewInboundRPCDispatch()

	// Create the server and pass in the dispatch class
	srvr := GeneratedDiscoveryServer{
		dispatch: d,
	}

	// Test calling the method TestCall
	errChan := make(chan error)
	defer close(errChan)
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	fn := func() {
		_, err := srvr.GetPeers(cancelCtx, &GetPeersRequest{})
		errChan <- err
	}
	go fn()
	cancelFunc()
	cancelErr := <-errChan
	assert.EqualError(t, cancelErr, "context canceled", "the error returned must be a context canceled error")
}

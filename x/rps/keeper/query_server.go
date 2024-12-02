package rpsKeeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"challenge/x/rps/types"
)

const (
	// Errors
	ErrEmptyRequest    = "empty request"
	ErrStudentNotFound = "student could not be found"
)

// Type assertion (interface)
var _ types.QueryServer = queryServer{}

// queryServer is the query server for the rps module
type queryServer struct {
	k Keeper
}

// NewQueryServerImpl returns a new query server
func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return queryServer{
		k: keeper,
	}
}

// GetStudent returns a single student by it's ID
func (qs queryServer) GetStudent(c context.Context, req *types.QueryGetStudentRequest) (*types.QueryGetStudentResponse, error) {
	// Wrap the context
	ctx := sdk.UnwrapSDKContext(c)

	// Check the request
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, ErrEmptyRequest)
	}

	// Check if the student can be found
	student, found := qs.k.GetStudent(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, ErrStudentNotFound)
	}

	// Return response
	return &types.QueryGetStudentResponse{Student: student}, nil
}

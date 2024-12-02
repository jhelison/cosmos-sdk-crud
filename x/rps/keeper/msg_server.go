package rpsKeeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrorstypes "github.com/cosmos/cosmos-sdk/types/errors"

	"challenge/x/rps/types"
)

const (
	// Errors
	ErrStudentExist = "Student with ID %s already exist"
)

// msgServer is the message server
type msgServer struct {
	k Keeper
}

// Type assertion (interface)
var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns a new message server
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{
		k: keeper,
	}
}

// CreateStudent create a new student using the message server
func (ms msgServer) CreateStudent(c context.Context, msg *types.MsgCreateStudent) (*types.MsgCreateStudentResponse, error) {
	// Wrap the context
	ctx := sdk.UnwrapSDKContext(c)

	// Check if the request is empty
	if msg == nil {
		return nil, sdkerrors.Wrap(sdkerrorstypes.ErrInvalidRequest, ErrEmptyRequest)
	}

	// Check if the student exists (for now we don't allow edits)
	_, found := ms.k.GetStudent(ctx, msg.Id)
	if found {
		return nil, sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidRequest, ErrStudentExist, msg.Id)
	}

	// Create the student
	student := types.NewStudent(msg.Name, msg.Id, msg.Age)

	// Validate
	if err := student.Validate(); err != nil {
		return nil, err
	}

	// Set the student
	if err := ms.k.SetStudent(ctx, student); err != nil {
		return nil, err
	}

	// Emit the event
	if err := ctx.EventManager().EmitTypedEvent(&types.EventNewStudent{
		Name: msg.Name,
		Id:   msg.Id,
		Age:  msg.Age,
	}); err != nil {
		return nil, err
	}

	// Return the response
	return &types.MsgCreateStudentResponse{}, nil
}

// DeleteStudent delete a student using the message server
func (ms msgServer) DeleteStudent(c context.Context, msg *types.MsgDeleteStudent) (*types.MsgDeleteStudentResponse, error) {
	// Wrap the context
	ctx := sdk.UnwrapSDKContext(c)

	// Check if the request is empty
	if msg == nil {
		return nil, sdkerrors.Wrap(sdkerrorstypes.ErrInvalidRequest, ErrEmptyRequest)
	}

	// Check if the student exists (it should exist)
	_, found := ms.k.GetStudent(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrorstypes.ErrInvalidRequest, ErrStudentNotFound)
	}

	// TODO: We may want to check for ownership

	// Delete the student
	if err := ms.k.RemoveStudent(ctx, msg.Id); err != nil {
		return nil, err
	}

	// Emit the event
	if err := ctx.EventManager().EmitTypedEvent(&types.EventDeleteStudent{Id: msg.Id}); err != nil {
		return nil, err
	}

	// Return the response
	return &types.MsgDeleteStudentResponse{}, nil
}

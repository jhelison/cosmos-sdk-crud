package rpsKeeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"challenge/x/rps/types"
)

// InitGenesis initializes the module state from a genesis state.
func (k *Keeper) InitGenesis(ctx sdk.Context, data *types.GenesisState) error {
	// Set the genesis games into the state
	for _, student := range data.Students {
		// Set the student
		err := k.SetStudent(ctx, student)
		if err != nil {
			return err
		}
	}

	return nil
}

// ExportGenesis exports the module state to a genesis state.
func (k *Keeper) ExportGenesis(ctx sdk.Context) (*types.GenesisState, error) {
	// Get all the students
	students, err := k.GetAllStudents(ctx)
	if err != nil {
		return nil, err
	}

	// Return the genesis
	return &types.GenesisState{
		Students: students,
	}, nil
}

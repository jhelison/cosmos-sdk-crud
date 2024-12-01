package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"challenge/x/rps/types"
)

// TestGenesisValidate tests the genesis validate
func TestGenesisValidate(t *testing.T) {
	studentId1 := sdk.AccAddress("student1")
	studentId2 := sdk.AccAddress("student2")

	// Test cases
	testCases := []struct {
		name        string
		genState    types.GenesisState
		errContains string
	}{
		{
			name:     "Valid - Default Genesis state",
			genState: *types.NewGenesisState(),
		},
		{
			name: "Valid - Genesis with valid entries",
			genState: types.GenesisState{
				[]types.Student{
					types.NewStudent("kii", studentId1.String(), 10),
					types.NewStudent("kii", studentId2.String(), 10),
				},
			},
		},
		{
			name: "Invalid - Duplicated entry",
			genState: types.GenesisState{
				[]types.Student{
					types.NewStudent("kii", studentId1.String(), 10),
					types.NewStudent("kii", studentId1.String(), 10),
				},
			},
			errContains: "id duplicated",
		},
		{
			name: "Valid - Invalid student",
			genState: types.GenesisState{
				[]types.Student{
					types.NewStudent("kii", studentId1.String(), 0), // Age zero
					types.NewStudent("kii", studentId2.String(), 10),
				},
			},
			errContains: "invalid age",
		},
	}

	// Run all the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Validate the genesis
			err := tc.genState.Validate()

			// Check if we should check against a error
			if tc.errContains != "" {
				require.ErrorContains(t, err, tc.errContains)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

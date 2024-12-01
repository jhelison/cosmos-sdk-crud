package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"challenge/x/rps/types"
)

// TestStudentValidate test the validate function from the student
// This all test all the getters
func TestStudentValidate(t *testing.T) {
	studentId := sdk.AccAddress("student")

	// Prepare the test cases
	testCases := []struct {
		name        string
		studentName string
		studentId   string
		studentAge  uint64
		errContains string
	}{
		{
			name:        "Pass - No error",
			studentName: "Kii",
			studentId:   studentId.String(),
			studentAge:  20,
		},
		{
			name:        "Error - address empty",
			studentName: "Kii",
			studentAge:  20,
			errContains: "empty address",
		},
		{
			name:        "Error - invalid age",
			studentName: "Kii",
			studentId:   studentId.String(),
			studentAge:  0,
			errContains: "invalid age",
		},
		{
			name:        "Error - empty name",
			studentName: "",
			studentId:   studentId.String(),
			studentAge:  20,
			errContains: "empty name",
		},
	}

	// Run all the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create the new student
			student := types.NewStudent(tc.studentName, tc.studentId, tc.studentAge)

			// Validate
			err := student.Validate()

			// Check if we should validate for error
			if tc.errContains != "" {
				require.ErrorContains(t, err, tc.errContains)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

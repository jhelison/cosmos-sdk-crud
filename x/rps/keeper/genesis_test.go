package rpsKeeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"challenge/x/rps/types"
)

// TestInitGenesis test the init genesis
func (s *KeeperTestSuite) TestInitGenesis() {
	// Define testing keys
	studentId1 := sdk.AccAddress("student1")
	studentId2 := sdk.AccAddress("student2")

	// Initialize the two students
	student1 := types.NewStudent("kii", studentId1.String(), 10)
	student2 := types.NewStudent("kii2", studentId2.String(), 20)

	// Create a new genesis state
	genesisState := types.GenesisState{
		Students: []types.Student{
			student1,
			student2,
		},
	}

	// Init the genesis
	err := s.k.InitGenesis(s.ctx, &genesisState)
	s.Require().NoError(err)

	// Check if the entries were created
	allStudents, err := s.k.GetAllStudents(s.ctx)
	s.Require().NoError(err)
	s.Require().Equal(allStudents, []types.Student{student2, student1})
}

// TestExportGenesis test the export genesis
func (s *KeeperTestSuite) TestExportGenesis() {
	// Ids for the students
	student1Key := sdk.AccAddress("student1")
	student2Key := sdk.AccAddress("student2")

	// Initialize the two students
	student1 := types.NewStudent("kii", student1Key.String(), 10)
	student2 := types.NewStudent("kii2", student2Key.String(), 20)

	// Store then on the KV store
	s.k.SetStudent(s.ctx, student1)
	s.k.SetStudent(s.ctx, student2)

	// Now run the export
	export, err := s.k.ExportGenesis(s.ctx)
	s.Require().NoError(err)
	s.Require().Equal(export.Students, []types.Student{student2, student1})
}

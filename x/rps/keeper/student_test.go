package rpsKeeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"challenge/x/rps/types"
)

// TestKeeperStudent tests the initialization of the keeper students KV access
func (s *KeeperTestSuite) TestKeeperStudent() {
	// Ids for the students
	student1Key := sdk.AccAddress("student1")
	student2Key := sdk.AccAddress("student2")

	// Initialize the two students
	student1 := types.NewStudent("kii", student1Key.String(), 10)
	student2 := types.NewStudent("kii2", student2Key.String(), 20)

	// Store then on the KV store
	s.k.SetStudent(s.ctx, student1)
	s.k.SetStudent(s.ctx, student2)

	// Look for the second student
	foundStudent, found := s.k.GetStudent(s.ctx, student2Key.String())
	s.Require().True(found)
	s.Require().Equal(foundStudent, student2)

	// Take all the students and check
	allStudents, err := s.k.GetAllStudents(s.ctx)
	s.Require().NoError(err)
	s.Require().Equal(allStudents, []types.Student{student2, student1})

	// Now delete the student 2 and check again
	err = s.k.RemoveStudent(s.ctx, student2Key.String())
	s.Require().NoError(err)

	// Try to fetch the student, should not be found
	_, found = s.k.GetStudent(s.ctx, student2Key.String())
	s.Require().False(found)

	// The first should still exist
	_, found = s.k.GetStudent(s.ctx, student1Key.String())
	s.Require().True(found)
}

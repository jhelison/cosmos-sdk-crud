package rpsKeeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"challenge/x/rps/types"
)

// TestMsgServerCreateStudent test the message server create student
func (s *KeeperTestSuite) TestMsgServerCreateStudent() {
	// Create a new student using the message server
	student1Key := sdk.MustAccAddressFromBech32("rps15pz3vytrun8apgnd8ka3zll6kkh9autyrs2kyj")

	// Try to create with no args
	_, err := s.msgServer.CreateStudent(s.ctx, nil)
	s.Require().ErrorContains(err, "empty request")

	// Run with bad args
	_, err = s.msgServer.CreateStudent(s.ctx, &types.MsgCreateStudent{Name: "kii", Id: student1Key.String(), Age: 0})
	s.Require().ErrorContains(err, "invalid age")

	// Now create with valid args
	response, err := s.msgServer.CreateStudent(s.ctx, &types.MsgCreateStudent{Name: "kii", Id: student1Key.String(), Age: 20})
	s.Require().NoError(err)
	s.Require().Equal(response, &types.MsgCreateStudentResponse{})

	// Now try to create with the same ID again
	_, err = s.msgServer.CreateStudent(s.ctx, &types.MsgCreateStudent{Name: "kii", Id: student1Key.String(), Age: 20})
	s.Require().ErrorContains(err, "already exist")
}

// TestMsgServerDeleteStudent test the delete student message
func (s *KeeperTestSuite) TestMsgServerDeleteStudent() {
	// Key for tests
	student1Key := sdk.MustAccAddressFromBech32("rps15pz3vytrun8apgnd8ka3zll6kkh9autyrs2kyj")

	// Try to delete with no args
	_, err := s.msgServer.DeleteStudent(s.ctx, nil)
	s.Require().ErrorContains(err, "empty request")

	// Try to delete a non existing key
	_, err = s.msgServer.DeleteStudent(s.ctx, &types.MsgDeleteStudent{Id: student1Key.String()})
	s.Require().ErrorContains(err, "student could not be found")

	// Create the student
	student1 := types.NewStudent("kii", student1Key.String(), 10)
	err = s.k.SetStudent(s.ctx, student1)
	s.Require().NoError(err)

	// Run a valid delete
	response, err := s.msgServer.DeleteStudent(s.ctx, &types.MsgDeleteStudent{Id: student1Key.String()})
	s.Require().NoError(err)
	s.Require().Equal(response, &types.MsgDeleteStudentResponse{})

	// The student now shouldn't exist
	_, found := s.k.GetStudent(s.ctx, student1Key.String())
	s.Require().False(found)
}

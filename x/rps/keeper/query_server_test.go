package rpsKeeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"challenge/x/rps/types"
)

// TestGRPCGetStudent test the get student query method
func (s *KeeperTestSuite) TestGRPCGetStudent() {
	// Try to send a empty arg
	_, err := s.queryServer.GetStudent(s.ctx, nil)
	s.Require().ErrorContains(err, "empty request")

	// Try to query a entry that doesn't exist
	_, err = s.queryServer.GetStudent(s.ctx, &types.QueryGetStudentRequest{
		Id: "student1",
	})
	s.Require().ErrorContains(err, "student could not be found")

	// Now register a student
	student1Key := sdk.AccAddress("student1")
	student := types.NewStudent("kii", student1Key.String(), 10)
	err = s.k.SetStudent(s.ctx, student)
	s.Require().NoError(err)

	// Try to get the new student
	response, err := s.queryServer.GetStudent(s.ctx, &types.QueryGetStudentRequest{
		Id: student1Key.String(),
	})
	s.Require().NoError(err)

	// Check the response
	s.Require().Equal(student, response.Student)

}

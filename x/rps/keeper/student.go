package rpsKeeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"challenge/x/rps/types"
)

// SetStudent set the student on on kv store
func (k Keeper) SetStudent(ctx sdk.Context, student types.Student) error {
	return k.Students.Set(ctx, student.Id, student)
}

// GetStudent returns a student from the kv store
func (k Keeper) GetStudent(ctx sdk.Context, id string) (student types.Student, found bool) {
	student, err := k.Students.Get(ctx, id)
	if err != nil {
		return student, false
	}
	return student, true
}

// RemoveStudent removes a student from the KV store
func (k Keeper) RemoveStudent(ctx sdk.Context, id string) error {
	return k.Students.Remove(ctx, id)
}

// GetAllStudents returns all the students
// This is used on the genesis export, if any other uses are needed we can do a iterator
func (k Keeper) GetAllStudents(ctx sdk.Context) (students []types.Student, err error) {
	// Walk the students
	err = k.Students.Walk(ctx, nil, func(id string, student types.Student) (stop bool, err error) {
		students = append(students, student)
		return false, nil
	})

	// If a error is found we return
	if err != nil {
		return nil, err
	}

	// Return the students
	return students, nil
}

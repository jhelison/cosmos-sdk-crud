package types

import (
	"errors"
	"strings"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Define the error messages
const (
	ErrEmptyAddress = "empty address"
	ErrInvalidAge   = "invalid age"
	ErrEmptyName    = "empty name"
)

// NewStudent returns a new student
func NewStudent(name string, id string, age uint64) Student {
	return Student{
		Name: name,
		Id:   id,
		Age:  age,
	}
}

// DefaultStudents return the default student
func DefaultStudents() (Students []Student) {
	return
}

// GetStudentAddress returns the parsed student ID
func (g Student) GetStudentAddress() (sdk.AccAddress, error) {
	return getStudentAddress(g.Id)
}

// GetStudentName returns the student name
func (g Student) GetStudentName() string {
	return g.Name
}

// GetStudentAge returns the student age
func (g Student) GetStudentAge() uint64 {
	return g.Age
}

// Validate validate the student object
// Validates if the address is empty, name is empty and if the age is less than 0
func (g Student) Validate() error {
	// Get Student id
	address, err := g.GetStudentAddress()
	if err != nil {
		return err
	}

	// Check if it's empty
	if address.Empty() {
		return errors.New(ErrEmptyAddress)
	}

	// This actually can't be reached
	if g.Age <= 0 {
		return errors.New(ErrInvalidAge)
	}

	// Validate the student name
	if len(strings.TrimSpace(g.Name)) == 0 {
		return errors.New(ErrEmptyName)
	}

	return nil
}

// getStudentAddress returns the parsed student address as acc address
func getStudentAddress(address string) (sdk.AccAddress, error) {
	// Validate the address has our prefix (it means the wallet is from out blockchain)
	addr, err := sdk.AccAddressFromBech32(address)
	return addr, sdkerrors.Wrap(err, "Invalid Address")
}

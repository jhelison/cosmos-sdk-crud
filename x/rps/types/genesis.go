package types

import "errors"

// Error messages
const (
	ErrIdDuplicated = "id duplicated"
)

// NewGenesisState creates a new genesis state with default values.
func NewGenesisState() *GenesisState {
	return &GenesisState{
		Students: DefaultStudents(), // Return the genesis value for a game
	}
}

// Validate performs basic genesis state validation returning an error upon any
func (gs *GenesisState) Validate() error {
	// Validate the genesis students
	unique := make(map[string]bool)
	for _, student := range gs.Students {
		// Validate if the student ID
		_, ok := unique[student.Id]
		if ok {
			return errors.New(ErrIdDuplicated)
		}

		// Set on the hashMap
		unique[student.Id] = true

		// Validate the student
		err := student.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

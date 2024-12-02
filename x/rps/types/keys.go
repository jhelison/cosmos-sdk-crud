package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "rps"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName
)

// Here store the stateless collection keys
var (
	StudentsKey = collections.NewPrefix(0)
)

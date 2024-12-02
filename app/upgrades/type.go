package upgrades

import (
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"challenge/app/keepers"
)

// Upgrade defines a chain upgrade
type Upgrade struct {
	// UpgradeName is the name of the upgrade (v1, v1_1, etc)
	UpgradeName string
	// CreateUpgradeHandler create the upgrade handler for the upgrade
	CreateUpgradeHandler func(*module.Manager, module.Configurator, *keepers.AppKeepers) upgradetypes.UpgradeHandler
}

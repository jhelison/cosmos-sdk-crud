package keepers

import (
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	consensuskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	rpsKeeper "challenge/x/rps/keeper"
)

// AppKeeper defines all the keepers for the app
type AppKeepers struct {
	// keepers
	AccountKeeper         authkeeper.AccountKeeper
	BankKeeper            bankkeeper.Keeper
	StakingKeeper         *stakingkeeper.Keeper
	DistrKeeper           distrkeeper.Keeper
	ConsensusParamsKeeper consensuskeeper.Keeper
	UpgradeKeeper         *upgradekeeper.Keeper
	RPSKeeper             rpsKeeper.Keeper // Import rps keeper
}

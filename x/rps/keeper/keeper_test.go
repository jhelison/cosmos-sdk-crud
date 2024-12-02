package rpsKeeper_test

import (
	"testing"
	"time"

	"cosmossdk.io/log"
	dbm "github.com/cosmos/cosmos-db"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"challenge/app"
	"challenge/app/params"
	rpsKeeper "challenge/x/rps/keeper"
	"challenge/x/rps/types"
)

// KeeperTestSuite is a suite with the keeper integration tests
type KeeperTestSuite struct {
	suite.Suite
	ctx sdk.Context
	app *app.RPSApp
	k   *rpsKeeper.Keeper

	// Query and message server
	queryServer types.QueryServer
}

// SetupTest sets up the testing environment before each test
func (suite *KeeperTestSuite) SetupTest() {
	// Set the bench32 prefix
	params.SetAddressPrefixes()

	// Prepare the params
	logger := log.NewNopLogger()
	db := dbm.NewMemDB()
	appOpts := simtestutil.NewAppOptionsWithFlagHome(suite.T().TempDir())

	// Start the app
	app, err := app.NewRPSApp(
		logger,
		db,
		nil,
		true,
		appOpts,
	)
	suite.Require().NoError(err)

	// Set the app on the suite
	suite.app = app

	// Start a new context
	suite.ctx = suite.app.NewContext(true)
	startTime := time.Now().UTC()
	suite.ctx = suite.ctx.WithBlockTime(startTime)

	// Set the keeper
	suite.k = &suite.app.RPSKeeper

	// Start the query server
	suite.queryServer = rpsKeeper.NewQueryServerImpl(*suite.k)
}

// TestKeeperTestSuite runs the full test suite
func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

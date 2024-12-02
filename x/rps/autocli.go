package rps

import (
	"fmt"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	rpsv1 "challenge/api/rps/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	// Get the bech32 address for the commands
	bech32PrefixAccAddr := sdk.GetConfig().GetBech32AccountAddrPrefix()

	// Return the CLI commands
	return &autocliv1.ModuleOptions{
		Tx: &autocliv1.ServiceCommandDescriptor{},
		// All the queries
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: rpsv1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				// The getStudent query
				{
					RpcMethod: "GetStudent",
					Use:       "student [id]",
					Short:     "Query a student by it's ID",
					Long: fmt.Sprintf(`Query a student by it's ID.
Example
$ %s query rps student %s1gghjut3ccd8ay0zduzj64hwre2fxs9ld75ru9p`, version.AppName, bech32PrefixAccAddr),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "id"},
					},
				},
			},
		},
	}
}

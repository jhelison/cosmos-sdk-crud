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
	// The app name
	appName := version.AppName

	// Return the CLI commands
	return &autocliv1.ModuleOptions{
		// All the txs
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: rpsv1.Msg_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				// The createStudent message
				{
					RpcMethod: "CreateStudent",
					Use:       "create_student [name] [id] [age]",
					Short:     "Create a new student based on a name, ID and age",
					Long: fmt.Sprintf(`Create a new student based on a name and age
Example
$ %s tx rps create_student kii 20`, bech32PrefixAccAddr),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "name"},
						{ProtoField: "id"},
						{ProtoField: "age"},
					},
				},
				// The deleteStudent message
				{
					RpcMethod: "DeleteStudent",
					Use:       "delete_student [id]",
					Short:     "Delete a student based on it's ID",
					Long: fmt.Sprintf(`Delete a student based on it's ID
Example
$ %s tx rps delete_student kii 20`, bech32PrefixAccAddr),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "id"},
					},
				},
			},
		},
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
$ %s query rps student %s1gghjut3ccd8ay0zduzj64hwre2fxs9ld75ru9p`, appName, bech32PrefixAccAddr),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "id"},
					},
				},
			},
		},
	}
}

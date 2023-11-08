package integration

import (
	appv1alpha1 "cosmossdk.io/api/cosmos/app/v1alpha1"
	authmodulev1 "cosmossdk.io/api/cosmos/auth/module/v1"
	"cosmossdk.io/core/appconfig"
	"github.com/cosmos/cosmos-sdk/testutil/configurator"
	disputemodulev1 "github.com/tellor-io/layer/api/layer/dispute/module"
	oraclemodulev1 "github.com/tellor-io/layer/api/layer/oracle/module"
	registrymodulev1 "github.com/tellor-io/layer/api/layer/registry/module"
)

func AuthModule() configurator.ModuleOption {
	return func(config *configurator.Config) {
		config.ModuleConfigs["auth"] = &appv1alpha1.ModuleConfig{
			Name: "auth",
			Config: appconfig.WrapAny(&authmodulev1.Module{
				Bech32Prefix: "cosmos",
				ModuleAccountPermissions: []*authmodulev1.ModuleAccountPermission{
					{Account: "fee_collector"},
					{Account: "distribution"},
					{Account: "oracle", Permissions: []string{"burner"}},
					{Account: "dispute"},
					{Account: "registry"},
					{Account: "mint", Permissions: []string{"minter"}},
					{Account: "bonded_tokens_pool", Permissions: []string{"burner", "staking"}},
					{Account: "not_bonded_tokens_pool", Permissions: []string{"burner", "staking"}},
					{Account: "gov", Permissions: []string{"burner"}},
					{Account: "nft"},
				},
			}),
		}
	}
}

func OracleModule() configurator.ModuleOption {
	return func(config *configurator.Config) {
		config.BeginBlockersOrder = append(config.BeginBlockersOrder, "oracle")
		config.EndBlockersOrder = append(config.EndBlockersOrder, "oracle")
		config.InitGenesisOrder = append(config.InitGenesisOrder, "oracle")
		config.ModuleConfigs["oracle"] = &appv1alpha1.ModuleConfig{
			Name:   "oracle",
			Config: appconfig.WrapAny(&oraclemodulev1.Module{}),
		}
	}
}

func DisputeModule() configurator.ModuleOption {
	return func(config *configurator.Config) {
		config.BeginBlockersOrder = append(config.BeginBlockersOrder, "dispute")
		config.EndBlockersOrder = append(config.EndBlockersOrder, "dispute")
		config.InitGenesisOrder = append(config.InitGenesisOrder, "dispute")
		config.ModuleConfigs["dispute"] = &appv1alpha1.ModuleConfig{
			Name:   "dispute",
			Config: appconfig.WrapAny(&disputemodulev1.Module{}),
		}
	}
}

func RegistryModule() configurator.ModuleOption {
	return func(config *configurator.Config) {
		config.BeginBlockersOrder = append(config.BeginBlockersOrder, "registry")
		config.EndBlockersOrder = append(config.EndBlockersOrder, "registry")
		config.InitGenesisOrder = append(config.InitGenesisOrder, "registry")
		config.ModuleConfigs["registry"] = &appv1alpha1.ModuleConfig{
			Name:   "registry",
			Config: appconfig.WrapAny(&registrymodulev1.Module{}),
		}
	}
}

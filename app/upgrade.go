package app

import (
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/gashub/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	bridgemoduletypes "github.com/bnb-chain/greenfield/x/bridge/types"
	paymentmodule "github.com/bnb-chain/greenfield/x/payment"
	paymenttypes "github.com/bnb-chain/greenfield/x/payment/types"
	storagemoduletypes "github.com/bnb-chain/greenfield/x/storage/types"
)

func (app *App) RegisterUpgradeHandlers(chainID string, serverCfg *serverconfig.Config) error {
	// Register the plans from server config
	err := app.UpgradeKeeper.RegisterUpgradePlan(chainID, serverCfg.Upgrade)
	if err != nil {
		return err
	}

	// Register the upgrade handlers here
	app.registerNagquUpgradeHandler()
	app.registerPampasUpgradeHandler()
	// app.register...()
	// ...
	return nil
}

// registerPublicDelegationUpgradeHandler registers the upgrade handlers for the public delegation upgrade.
// it will be enabled at the future version.
// func (app *App) registerPublicDelegationUpgradeHandler() {
// 	// Register the upgrade handler
// 	app.UpgradeKeeper.SetUpgradeHandler(upgradetypes.EnablePublicDelegationUpgrade,
// 		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
// 			app.Logger().Info("upgrade to ", plan.Name)
// 			return fromVM, nil
// 		})

// 	// Register the upgrade initializer
// 	app.UpgradeKeeper.SetUpgradeInitializer(upgradetypes.EnablePublicDelegationUpgrade,
// 		func() error {
// 			app.Logger().Info("Init enable public delegation upgrade")
// 			return nil
// 		},
// 	)
// }

func (app *App) registerNagquUpgradeHandler() {
	// Register the upgrade handler
	app.UpgradeKeeper.SetUpgradeHandler(upgradetypes.Nagqu,
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			app.Logger().Info("upgrade to ", plan.Name)
			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		})

	// Register the upgrade initializer
	app.UpgradeKeeper.SetUpgradeInitializer(upgradetypes.Nagqu,
		func() error {
			app.Logger().Info("Init Nagqu upgrade")
			mm, ok := app.mm.Modules[paymenttypes.ModuleName].(*paymentmodule.AppModule)
			if !ok {
				panic("*paymentmodule.AppModule not found")
			}
			mm.SetConsensusVersion(2)
			return nil
		})
}

func (app *App) registerPampasUpgradeHandler() {
	// Register the upgrade handler
	app.UpgradeKeeper.SetUpgradeHandler(upgradetypes.Pampas,
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			app.Logger().Info("upgrade to ", plan.Name)

			// open resource channels for opbnb
			app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestOpChainId), bridgemoduletypes.SyncParamsChannelID, sdk.ChannelAllow)
			app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestOpChainId), storagemoduletypes.BucketChannelId, sdk.ChannelAllow)
			app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestOpChainId), storagemoduletypes.ObjectChannelId, sdk.ChannelAllow)
			app.CrossChainKeeper.SetChannelSendPermission(ctx, sdk.ChainID(app.appConfig.CrossChain.DestOpChainId), storagemoduletypes.GroupChannelId, sdk.ChannelAllow)

			// register MsgRejectMigrateBucket Gas param
			app.GashubKeeper.SetMsgGasParams(ctx, *types.NewMsgGasParamsWithFixedGas("/greenfield.storage.MsgRejectMigrateBucket", 1.2e3))

			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		})

	// Register the upgrade initializer
	app.UpgradeKeeper.SetUpgradeInitializer(upgradetypes.Pampas,
		func() error {
			app.Logger().Info("Init Pampas upgrade")

			// enable chain id for opbnb
			app.CrossChainKeeper.SetDestOpChainID(sdk.ChainID(app.appConfig.CrossChain.DestOpChainId))
			return nil
		})
}

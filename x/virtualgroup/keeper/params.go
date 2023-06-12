package keeper

import (
	"cosmossdk.io/math"
	"github.com/bnb-chain/greenfield/x/virtualgroup/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) DepositDenomForGVG(ctx sdk.Context) (res string) {
	params := k.GetParams(ctx)
	return params.DepositDenom
}

func (k Keeper) MinDeposit(ctx sdk.Context) (res math.Int) {
	params := k.GetParams(ctx)
	return params.MinDeposit
}

func (k Keeper) GVGStakingPrice(ctx sdk.Context) (res sdk.Dec) {
	params := k.GetParams(ctx)
	return params.GvgStakingPrice
}

// GetParams returns the current sp module parameters.
func (k Keeper) GetParams(ctx sdk.Context) (p types.Params) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ParamsKey)
	if bz == nil {
		return p
	}

	k.cdc.MustUnmarshal(bz, &p)
	return p
}

// SetParams sets the params of sp module
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	if err := params.Validate(); err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&params)
	store.Set(types.ParamsKey, bz)

	return nil
}

package keeper

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bnb-chain/greenfield/x/payment/types"
)

func (k Keeper) CheckStreamRecord(streamRecord *types.StreamRecord) {
	if streamRecord == nil {
		panic("streamRecord is nil")
	}
	if len(streamRecord.Account) != sdk.EthAddressLength*2+2 {
		panic(fmt.Sprintf("invalid streamRecord account %s", streamRecord.Account))
	}
	if streamRecord.Status != types.STREAM_ACCOUNT_STATUS_ACTIVE && streamRecord.Status != types.STREAM_ACCOUNT_STATUS_FROZEN {
		panic(fmt.Sprintf("invalid streamRecord status %d", streamRecord.Status))
	}
	if streamRecord.StaticBalance.IsNil() {
		panic(fmt.Sprintf("invalid streamRecord staticBalance %s", streamRecord.StaticBalance))
	}
	if streamRecord.NetflowRate.IsNil() {
		panic(fmt.Sprintf("invalid streamRecord netflowRate %s", streamRecord.NetflowRate))
	}
	if streamRecord.LockBalance.IsNil() || streamRecord.LockBalance.IsNegative() {
		panic(fmt.Sprintf("invalid streamRecord lockBalance %s", streamRecord.LockBalance))
	}
	if streamRecord.BufferBalance.IsNil() || streamRecord.BufferBalance.IsNegative() {
		panic(fmt.Sprintf("invalid streamRecord bufferBalance %s", streamRecord.BufferBalance))
	}
}

// SetStreamRecord set a specific streamRecord in the store from its index
func (k Keeper) SetStreamRecord(ctx sdk.Context, streamRecord *types.StreamRecord) {
	k.CheckStreamRecord(streamRecord)
	account := streamRecord.Account
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StreamRecordKeyPrefix)
	key := types.StreamRecordKey(sdk.MustAccAddressFromHex(account))
	streamRecord.Account = ""
	b := k.cdc.MustMarshal(streamRecord)
	store.Set(key, b)
	// set the field back, the streamRecord may be used after this function
	streamRecord.Account = account
	event := &types.EventStreamRecordUpdate{
		Account:           streamRecord.Account,
		StaticBalance:     streamRecord.StaticBalance,
		NetflowRate:       streamRecord.NetflowRate,
		FrozenNetflowRate: streamRecord.FrozenNetflowRate,
		CrudTimestamp:     streamRecord.CrudTimestamp,
		Status:            streamRecord.Status,
		LockBalance:       streamRecord.LockBalance,
		BufferBalance:     streamRecord.BufferBalance,
		SettleTimestamp:   streamRecord.SettleTimestamp,
	}
	_ = ctx.EventManager().EmitTypedEvents(event)
}

// GetStreamRecord returns a streamRecord from its index
func (k Keeper) GetStreamRecord(
	ctx sdk.Context,
	account sdk.AccAddress,
) (val *types.StreamRecord, found bool) {
	val = types.NewStreamRecord(account, ctx.BlockTime().Unix())
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StreamRecordKeyPrefix)

	b := store.Get(types.StreamRecordKey(
		account,
	))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, val)
	val.Account = account.String()
	return val, true
}

// GetAllStreamRecord returns all streamRecord
func (k Keeper) GetAllStreamRecord(ctx sdk.Context) (list []types.StreamRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StreamRecordKeyPrefix)
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StreamRecord
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		val.Account = string(iterator.Key())
		list = append(list, val)
	}

	return
}

// UpdateFrozenStreamRecord updates frozen streamRecord in `force delete` scenarios
// it only handles the lock balance change and ignore the other changes(since the streams are already changed and the
// accumulated OutFlows are changed outside this function)
func (k Keeper) UpdateFrozenStreamRecord(ctx sdk.Context, streamRecord *types.StreamRecord, change *types.StreamRecordChange) error {
	if streamRecord.Status != types.STREAM_ACCOUNT_STATUS_FROZEN {
		return fmt.Errorf("stream account %s is not frozen", streamRecord.Account)
	}
	currentTimestamp := ctx.BlockTime().Unix()
	streamRecord.CrudTimestamp = currentTimestamp
	// update lock balance
	if !change.LockBalanceChange.IsZero() {
		streamRecord.LockBalance = streamRecord.LockBalance.Add(change.LockBalanceChange)
		streamRecord.StaticBalance = streamRecord.StaticBalance.Sub(change.LockBalanceChange)
		if streamRecord.LockBalance.IsNegative() {
			return fmt.Errorf("lock balance can not become negative, current: %s", streamRecord.LockBalance)
		}
	}
	return nil
}

func (k Keeper) UpdateStreamRecord(ctx sdk.Context, streamRecord *types.StreamRecord, change *types.StreamRecordChange) error {
	if streamRecord.Status != types.STREAM_ACCOUNT_STATUS_ACTIVE {
		if streamRecord.Status == types.STREAM_ACCOUNT_STATUS_FROZEN {
			if forced, ok := ctx.Value(types.ForceUpdateFrozenStreamRecordKey).(bool); forced && ok {
				return k.UpdateFrozenStreamRecord(ctx, streamRecord, change)
			}
		}
		return fmt.Errorf("stream account %s is frozen", streamRecord.Account)
	}
	isPay := change.StaticBalanceChange.IsNegative() || change.RateChange.IsNegative()
	currentTimestamp := ctx.BlockTime().Unix()
	timestamp := streamRecord.CrudTimestamp
	params := k.GetParams(ctx)
	// update delta balance
	if currentTimestamp != timestamp {
		if !streamRecord.NetflowRate.IsZero() {
			flowDelta := streamRecord.NetflowRate.MulRaw(currentTimestamp - timestamp)
			streamRecord.StaticBalance = streamRecord.StaticBalance.Add(flowDelta)
		}
		streamRecord.CrudTimestamp = currentTimestamp
	}
	// update lock balance
	if !change.LockBalanceChange.IsZero() {
		streamRecord.LockBalance = streamRecord.LockBalance.Add(change.LockBalanceChange)
		streamRecord.StaticBalance = streamRecord.StaticBalance.Sub(change.LockBalanceChange)
		if streamRecord.LockBalance.IsNegative() {
			return fmt.Errorf("lock balance can not become negative, current: %s", streamRecord.LockBalance)
		}
	}
	// update buffer balance
	if !change.RateChange.IsZero() {
		streamRecord.NetflowRate = streamRecord.NetflowRate.Add(change.RateChange)
		newBufferBalance := sdkmath.ZeroInt()
		if streamRecord.NetflowRate.IsNegative() {
			newBufferBalance = streamRecord.NetflowRate.Abs().Mul(sdkmath.NewIntFromUint64(params.VersionedParams.ReserveTime))
		}
		if !newBufferBalance.Equal(streamRecord.BufferBalance) {
			streamRecord.StaticBalance = streamRecord.StaticBalance.Sub(newBufferBalance).Add(streamRecord.BufferBalance)
			streamRecord.BufferBalance = newBufferBalance
		}
	}
	// update static balance
	if !change.StaticBalanceChange.IsZero() {
		streamRecord.StaticBalance = streamRecord.StaticBalance.Add(change.StaticBalanceChange)
	}
	if streamRecord.StaticBalance.IsNegative() {
		account := sdk.MustAccAddressFromHex(streamRecord.Account)
		hasBankAccount := k.accountKeeper.HasAccount(ctx, account)
		if hasBankAccount {
			coins := sdk.NewCoins(sdk.NewCoin(params.FeeDenom, streamRecord.StaticBalance.Abs()))
			err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, account, types.ModuleName, coins)
			if err != nil {
				ctx.Logger().Info("auto transfer failed", "account", streamRecord.Account, "err", err, "coins", coins)
			} else {
				streamRecord.StaticBalance = sdkmath.ZeroInt()
			}
		}
	}
	// if the change is a pay(which decreases the static balance or netflow rate), the left static balance should be enough
	if isPay && streamRecord.StaticBalance.IsNegative() {
		return fmt.Errorf("stream account %s balance not enough, lack of %s BNB wei", streamRecord.Account, streamRecord.StaticBalance.Abs())
	}
	//calculate settle time
	var settleTimestamp int64 = 0
	if streamRecord.NetflowRate.IsNegative() {
		payDuration := streamRecord.StaticBalance.Add(streamRecord.BufferBalance).Quo(streamRecord.NetflowRate.Abs())
		if payDuration.LTE(sdkmath.NewIntFromUint64(params.ForcedSettleTime)) {
			return fmt.Errorf("stream account %s balance not enough, lack of %s BNB", streamRecord.Account, streamRecord.StaticBalance.Abs())
		} else {
			settleTimestamp = currentTimestamp - int64(params.ForcedSettleTime) + payDuration.Int64()
		}
	}
	k.UpdateAutoSettleRecord(ctx, sdk.MustAccAddressFromHex(streamRecord.Account), streamRecord.SettleTimestamp, settleTimestamp)
	streamRecord.SettleTimestamp = settleTimestamp
	return nil
}

func (k Keeper) SettleStreamRecord(ctx sdk.Context, streamRecord *types.StreamRecord) error {
	currentTimestamp := ctx.BlockTime().Unix()
	timestamp := streamRecord.CrudTimestamp
	params := k.GetParams(ctx)

	if currentTimestamp != timestamp {
		if !streamRecord.NetflowRate.IsZero() {
			flowDelta := streamRecord.NetflowRate.MulRaw(currentTimestamp - timestamp)
			streamRecord.StaticBalance = streamRecord.StaticBalance.Add(flowDelta)
		}
		streamRecord.CrudTimestamp = currentTimestamp
	}

	if streamRecord.StaticBalance.IsNegative() {
		account := sdk.MustAccAddressFromHex(streamRecord.Account)
		hasBankAccount := k.accountKeeper.HasAccount(ctx, account)
		if hasBankAccount {
			coins := sdk.NewCoins(sdk.NewCoin(params.FeeDenom, streamRecord.StaticBalance.Abs()))
			err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, account, types.ModuleName, coins)
			if err != nil {
				ctx.Logger().Info("auto transfer failed", "account", streamRecord.Account, "err", err, "coins", coins)
			} else {
				streamRecord.StaticBalance = sdkmath.ZeroInt()
			}
		}
	}

	var settleTimestamp int64 = 0
	if streamRecord.NetflowRate.IsNegative() {
		payDuration := streamRecord.StaticBalance.Add(streamRecord.BufferBalance).Quo(streamRecord.NetflowRate.Abs())
		if payDuration.LTE(sdkmath.NewIntFromUint64(params.ForcedSettleTime)) {
			err := k.ForceSettle(ctx, streamRecord)
			if err != nil {
				return err
			}
		} else {
			settleTimestamp = currentTimestamp - int64(params.ForcedSettleTime) + payDuration.Int64()
		}
	}
	k.UpdateAutoSettleRecord(ctx, sdk.MustAccAddressFromHex(streamRecord.Account), streamRecord.SettleTimestamp, settleTimestamp)
	streamRecord.SettleTimestamp = settleTimestamp
	return nil
}

func (k Keeper) UpdateStreamRecordByAddr(ctx sdk.Context, change *types.StreamRecordChange) (ret *types.StreamRecord, err error) {
	streamRecord, _ := k.GetStreamRecord(ctx, change.Addr)
	err = k.UpdateStreamRecord(ctx, streamRecord, change)
	if err != nil {
		return
	}
	k.SetStreamRecord(ctx, streamRecord)
	return streamRecord, nil
}

func (k Keeper) ForceSettle(ctx sdk.Context, streamRecord *types.StreamRecord) error {
	totalBalance := streamRecord.StaticBalance.Add(streamRecord.BufferBalance)
	change := types.NewDefaultStreamRecordChangeWithAddr(types.GovernanceAddress).WithStaticBalanceChange(totalBalance)
	_, err := k.UpdateStreamRecordByAddr(ctx, change)
	if err != nil {
		return fmt.Errorf("update governance stream record failed: %w", err)
	}
	// force settle
	streamRecord.StaticBalance = sdkmath.ZeroInt()
	streamRecord.BufferBalance = sdkmath.ZeroInt()
	streamRecord.Status = types.STREAM_ACCOUNT_STATUS_FROZEN
	// emit event
	_ = ctx.EventManager().EmitTypedEvents(&types.EventForceSettle{
		Addr:           streamRecord.Account,
		SettledBalance: totalBalance,
	})
	return nil
}

func (k Keeper) AutoSettle(ctx sdk.Context) {
	currentTimestamp := ctx.BlockTime().Unix()
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AutoSettleRecordKeyPrefix)
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	count := uint64(0)
	max := k.GetParams(ctx).MaxAutoSettleFlowCount
	for ; iterator.Valid(); iterator.Next() {
		if count >= max {
			return
		}
		val := types.ParseAutoSettleRecordKey(iterator.Key())
		addr := sdk.MustAccAddressFromHex(val.Addr)
		if val.Timestamp > currentTimestamp {
			return
		}
		streamRecord, found := k.GetStreamRecord(ctx, addr)
		if !found {
			ctx.Logger().Error("stream record not found", "addr", val.Addr)
			panic("stream record not found")
		}

		if streamRecord.Status == types.STREAM_ACCOUNT_STATUS_ACTIVE {
			err := k.SettleStreamRecord(ctx, streamRecord)
			if err != nil {
				panic(err)
			}
			count++ // add one for a stream record
			if streamRecord.Status == types.STREAM_ACCOUNT_STATUS_ACTIVE {
				k.SetStreamRecord(ctx, streamRecord)
				continue
			}
		}

		activeFlowKey := types.OutFlowKey(addr, types.OUT_FLOW_STATUS_ACTIVE, nil)
		flowStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.OutFlowKeyPrefix)
		flowIterator := flowStore.Iterator(activeFlowKey, nil)
		defer flowIterator.Close()

		totalRate := sdk.ZeroInt()
		for ; flowIterator.Valid(); flowIterator.Next() {
			if count >= max {
				break
			}
			_, outFlow := types.ParseOutFlowKey(flowIterator.Key())
			if outFlow.Status == types.OUT_FLOW_STATUS_FROZEN {
				break
			}
			rate := types.ParseOutFlowValue(flowIterator.Value())

			toAddr := sdk.MustAccAddressFromHex(outFlow.ToAddress)
			flowChange := types.NewDefaultStreamRecordChangeWithAddr(toAddr).WithRateChange(rate)
			_, err := k.UpdateStreamRecordByAddr(ctx, flowChange)
			if err != nil {
				panic(fmt.Sprintf("update %s stream record failed: %s", outFlow.ToAddress, err.Error()))
			}

			flowStore.Delete(flowIterator.Key())
			outFlow.Status = types.OUT_FLOW_STATUS_FROZEN
			k.SetOutFlow(ctx, addr, &outFlow)

			totalRate = totalRate.Add(rate)
			count++
		}
		streamRecord.NetflowRate = streamRecord.NetflowRate.Add(totalRate)
		streamRecord.FrozenNetflowRate = totalRate.Add(streamRecord.FrozenNetflowRate)
		k.SetStreamRecord(ctx, streamRecord)
	}
}

func (k Keeper) TryResumeStreamRecord(ctx sdk.Context, streamRecord *types.StreamRecord, depositBalance sdkmath.Int) error {
	if streamRecord.Status != types.STREAM_ACCOUNT_STATUS_FROZEN {
		return fmt.Errorf("stream account %s status is not frozen", streamRecord.Account)
	}

	if !streamRecord.NetflowRate.IsZero() { // the account is resuming
		return fmt.Errorf("stream account %s status is resuming, although it is frozen now", streamRecord.Account)
	}

	params := k.GetParams(ctx)
	reserveTime := params.VersionedParams.ReserveTime
	forcedSettleTime := params.ForcedSettleTime

	totalRate := streamRecord.NetflowRate.Add(streamRecord.FrozenNetflowRate)
	streamRecord.StaticBalance = streamRecord.StaticBalance.Add(depositBalance)
	expectedBalanceToResume := totalRate.Mul(sdkmath.NewIntFromUint64(reserveTime))
	if streamRecord.StaticBalance.LT(expectedBalanceToResume) {
		// deposit balance is not enough to resume, only add static balance
		k.SetStreamRecord(ctx, streamRecord)
		return nil
	}

	now := ctx.BlockTime().Unix()
	ctx.Logger().Debug("try to resume stream account", "streamRecord.OutFlowCount", streamRecord.OutFlowCount, "params.MaxAutoResumeFlowCount", params.MaxAutoResumeFlowCount)
	if streamRecord.OutFlowCount <= params.MaxAutoResumeFlowCount { //only rough judgement, resume directly
		streamRecord.Status = types.STREAM_ACCOUNT_STATUS_ACTIVE
		streamRecord.SettleTimestamp = now + streamRecord.StaticBalance.Quo(totalRate).Int64() - int64(forcedSettleTime)
		streamRecord.NetflowRate = totalRate
		streamRecord.FrozenNetflowRate = sdkmath.ZeroInt()
		streamRecord.BufferBalance = expectedBalanceToResume
		streamRecord.StaticBalance = streamRecord.StaticBalance.Sub(expectedBalanceToResume)
		streamRecord.CrudTimestamp = now

		addr := sdk.MustAccAddressFromHex(streamRecord.Account)
		frozenFlowKey := types.OutFlowKey(addr, types.OUT_FLOW_STATUS_FROZEN, nil)
		flowStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.OutFlowKeyPrefix)
		flowIterator := flowStore.Iterator(frozenFlowKey, nil)
		defer flowIterator.Close()

		for ; flowIterator.Valid(); flowIterator.Next() {
			_, outFlow := types.ParseOutFlowKey(flowIterator.Key())
			rate := types.ParseOutFlowValue(flowIterator.Value())

			toAddr := sdk.MustAccAddressFromHex(outFlow.ToAddress)
			change := types.NewDefaultStreamRecordChangeWithAddr(toAddr).WithRateChange(rate)
			_, err := k.UpdateStreamRecordByAddr(ctx, change)
			if err != nil {
				return fmt.Errorf("update receiver stream record failed: %w", err)
			}

			flowStore.Delete(flowIterator.Key())
			outFlow.Status = types.OUT_FLOW_STATUS_ACTIVE
			k.SetOutFlow(ctx, addr, &outFlow)
		}
		k.SetStreamRecord(ctx, streamRecord)
		k.UpdateAutoSettleRecord(ctx, sdk.MustAccAddressFromHex(streamRecord.Account), 0, streamRecord.SettleTimestamp)
		return nil
	} else { //enqueue for resume in end block
		k.SetAutoResumeRecord(ctx, &types.AutoResumeRecord{
			Timestamp: now,
			Addr:      streamRecord.Account,
		})
		return nil
	}
}

func (k Keeper) AutoResume(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AutoResumeRecordKeyPrefix)
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	var count uint64 = 0
	max := k.GetParams(ctx).MaxAutoResumeFlowCount
	for ; iterator.Valid(); iterator.Next() {
		autoResumeRecord := types.ParseAutoResumeRecordKey(iterator.Key())
		addr := sdk.MustAccAddressFromHex(autoResumeRecord.Addr)

		streamRecord, found := k.GetStreamRecord(ctx, addr)
		if !found {
			ctx.Logger().Error("stream record not found", "addr", autoResumeRecord.Addr)
			panic("stream record not found")
		}

		totalRate := sdk.ZeroInt()
		frozenFlowKey := types.OutFlowKey(addr, types.OUT_FLOW_STATUS_FROZEN, nil)
		flowStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.OutFlowKeyPrefix)
		flowIterator := flowStore.Iterator(frozenFlowKey, nil)
		defer flowIterator.Close()

		for ; flowIterator.Valid(); flowIterator.Next() {
			if count >= max {
				break
			}
			_, outFlow := types.ParseOutFlowKey(flowIterator.Key())
			rate := types.ParseOutFlowValue(flowIterator.Value())

			toAddr := sdk.MustAccAddressFromHex(outFlow.ToAddress)
			flowChange := types.NewDefaultStreamRecordChangeWithAddr(toAddr).WithRateChange(rate)
			_, err := k.UpdateStreamRecordByAddr(ctx, flowChange)
			if err != nil {
				panic(fmt.Sprintf("update %s stream record failed: %s", outFlow.ToAddress, err.Error()))
			}

			flowStore.Delete(flowIterator.Key())
			outFlow.Status = types.OUT_FLOW_STATUS_ACTIVE
			k.SetOutFlow(ctx, addr, &outFlow)

			totalRate = totalRate.Add(rate)
			count++
		}

		streamRecord.NetflowRate = streamRecord.NetflowRate.Add(totalRate.Neg())
		streamRecord.FrozenNetflowRate = streamRecord.NetflowRate.Add(totalRate)
		if !flowIterator.Valid() {
			if !streamRecord.FrozenNetflowRate.IsZero() {
				panic("should not happen") // TODO: assertion for fail quick, remove later
			}
			streamRecord.Status = types.STREAM_ACCOUNT_STATUS_ACTIVE
			change := types.NewDefaultStreamRecordChangeWithAddr(addr)
			_, err := k.UpdateStreamRecordByAddr(ctx, change)
			if err != nil {
				panic(fmt.Sprintf("update %s stream record failed: %s", addr, err.Error()))
			}
			k.RemoveAutoResumeRecord(ctx, autoResumeRecord.Timestamp, addr)
		}
		k.SetStreamRecord(ctx, streamRecord)
	}
}

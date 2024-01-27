package keeper

import (
	"encoding/hex"
	"fmt"

	"cosmossdk.io/math"
	"cosmossdk.io/store"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tellor-io/layer/lib"
	"github.com/tellor-io/layer/x/oracle/types"
)

func (k Keeper) transfer(ctx sdk.Context, tipper sdk.AccAddress, tip sdk.Coin) (sdk.Coin, error) {
	twoPercent := tip.Amount.Mul(math.NewInt(2)).Quo(math.NewInt(100))
	burnCoin := sdk.NewCoin(tip.Denom, twoPercent)
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, tipper, types.ModuleName, sdk.NewCoins(tip)); err != nil {
		return sdk.NewCoin(tip.Denom, math.ZeroInt()), err
	}
	// burn 2% of tip
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(burnCoin)); err != nil {
		return sdk.NewCoin(tip.Denom, math.ZeroInt()), err
	}
	return tip.Sub(burnCoin), nil
}

func (k Keeper) SetTip(ctx sdk.Context, tipper sdk.AccAddress, queryData string, tip sdk.Coin) {
	tipStore := k.TipStore(ctx)
	k.SetTotalTips(ctx, tipStore, tip)
	k.SetQueryTips(ctx, tipStore, queryData, tip)
	k.SetTipperTipsForQuery(ctx, tipStore, tipper.String(), queryData, tip)
	k.SetTipperTotalTips(ctx, tipStore, tipper, tip)
}

func (k Keeper) SetQueryTips(ctx sdk.Context, tipStore storetypes.KVStore, queryData string, tip sdk.Coin) {
	tips, queryId := k.GetQueryTips(ctx, tipStore, queryData)
	tips.Amount = tips.Amount.Add(tip)
	tips.TotalTips = tips.TotalTips.Add(tip)
	tipStore.Set(queryId, k.cdc.MustMarshal(&tips))
}

func (k Keeper) SetTipperTipsForQuery(ctx sdk.Context, tipStore store.KVStore, tipper, queryData string, tip sdk.Coin) {
	tips := k.GetUserQueryTips(ctx, tipper, queryData)
	tips.Total = tips.Total.Add(tip)
	tipStore.Set(k.TipperKey(tipper, queryData), k.cdc.MustMarshal(&tips))
}

func (k Keeper) SetTipperTotalTips(ctx sdk.Context, tipStore store.KVStore, tipper sdk.AccAddress, tip sdk.Coin) {
	tips := k.GetUserTips(ctx, tipper)
	tips.Total = tips.Total.Add(tip)
	tipStore.Set(tipper, k.cdc.MustMarshal(&tips))
}

func (k Keeper) SetTotalTips(ctx sdk.Context, tipStore store.KVStore, tip sdk.Coin) {
	total := k.GetTotalTips(ctx)
	totals := total.Add(tip)
	tipStore.Set([]byte("totaltips"), k.cdc.MustMarshal(&totals))
}

func (k Keeper) GetQueryTips(ctx sdk.Context, tipStore store.KVStore, queryData string) (types.Tips, []byte) {
	if lib.Has0xPrefix(queryData) {
		queryData = queryData[2:]
	}
	// decode query data hex string to bytes
	queryDataBytes, err := hex.DecodeString(queryData)
	if err != nil {
		panic(err)
	}
	queryId := HashQueryData(queryDataBytes)
	bz := tipStore.Get(queryId)
	if bz == nil {
		return types.Tips{
			QueryData: queryData,
			Amount:    sdk.NewCoin(types.DefaultBondDenom, math.ZeroInt()),
			TotalTips: sdk.NewCoin(types.DefaultBondDenom, math.ZeroInt()),
		}, queryId
	}
	var tips types.Tips
	k.cdc.Unmarshal(bz, &tips)
	return tips, queryId
}

func (k Keeper) GetQueryTip(ctx sdk.Context, queryId string) sdk.Coin {
	tipStore := k.TipStore(ctx)
	qId, _ := hex.DecodeString(queryId)
	bz := tipStore.Get(qId)
	if bz == nil {
		return sdk.NewCoin(types.DefaultBondDenom, math.ZeroInt())
	}
	var tips types.Tips
	k.cdc.Unmarshal(bz, &tips)
	return tips.Amount
}

func (k Keeper) GetUserTips(ctx sdk.Context, tipper sdk.AccAddress) types.UserTipTotal {
	tipStore := k.TipStore(ctx)
	bz := tipStore.Get(tipper)
	if bz == nil {
		return types.UserTipTotal{
			Address: tipper.String(),
			Total:   sdk.NewCoin(types.DefaultBondDenom, math.ZeroInt()),
		}
	}
	var tips types.UserTipTotal
	k.cdc.Unmarshal(bz, &tips)
	return tips
}

func (k Keeper) GetUserQueryTips(ctx sdk.Context, tipper, queryData string) (tips types.UserTipTotal) {
	tipStore := k.TipStore(ctx)
	if lib.Has0xPrefix(queryData) {
		queryData = queryData[2:]
	}
	bz := tipStore.Get(k.TipperKey(tipper, queryData))
	if bz == nil {
		return types.UserTipTotal{
			Address: tipper,
			Total:   sdk.NewCoin(types.DefaultBondDenom, math.ZeroInt()),
		}
	}
	k.cdc.Unmarshal(bz, &tips)
	return
}

func (k Keeper) TipperKey(tipper, queryData string) []byte {
	return types.KeyPrefix(fmt.Sprintf("%s:%s", tipper, queryData))
}

func (k Keeper) GetTotalTips(ctx sdk.Context) sdk.Coin {
	tipStore := k.TipStore(ctx)
	bz := tipStore.Get([]byte("totaltips"))
	if bz == nil {
		return sdk.NewCoin(types.DefaultBondDenom, math.ZeroInt())
	}
	var total sdk.Coin
	k.cdc.MustUnmarshal(bz, &total)
	return total
}

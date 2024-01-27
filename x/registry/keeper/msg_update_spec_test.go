package keeper_test

import (
	"testing"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/stretchr/testify/require"
	"github.com/tellor-io/layer/x/registry/types"
)

func TestUpdateDataSpec(t *testing.T) {
	ms, ctx, k := setupMsgServer(t)
	authority := authtypes.NewModuleAddress(govtypes.ModuleName).String()
	queryType := "testQueryType"
	spec := types.DataSpec{
		DocumentHash:      "testHash",
		ValueType:         "uint256",
		AggregationMethod: "weighted-median",
	}

	// Register spec
	registerSpecInput := &types.MsgRegisterSpec{
		Registrar: "creator1",
		QueryType: queryType,
		Spec:      spec,
	}
	_, err := ms.RegisterSpec(ctx, registerSpecInput)
	require.NoError(t, err)
	// Update spec
	spec.ValueType = "uint128"
	updateSpecInput := &types.MsgUpdateDataSpec{
		Authority: authority,
		QueryType: queryType,
		Spec:      spec,
	}
	_, err = ms.UpdateDataSpec(ctx, updateSpecInput)
	require.NoError(t, err)

	// Check if spec is updated
	getSpec, err := k.SpecRegistry.Get(ctx, queryType)
	require.NoError(t, err)
	require.Equal(t, "uint128", getSpec.ValueType)

	// Update spec with invalid authority
	authority = "invalidAuthority"
	// update spec
	spec.ValueType = "int256"
	updateSpecInput = &types.MsgUpdateDataSpec{
		Authority: authority,
		QueryType: queryType,
		Spec:      spec,
	}
	_, err = ms.UpdateDataSpec(ctx, updateSpecInput)
	require.ErrorContains(t, err, "invalidAuthority")
}

package keeper

import (
	"context"
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tellor-io/layer/x/registry/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tellor-io/layer/lib"
)

func (k Querier) DecodeQuerydata(goCtx context.Context, req *types.QueryDecodeQuerydataRequest) (*types.QueryDecodeQuerydataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// remove 0x from hex string if present
	req.Querydata = lib.RemoveHexPrefix(req.Querydata)
	// decode query data hex to bytes
	queryDataBytes, err := hex.DecodeString(req.Querydata)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("failed to decode query data string: %v", err))
	}
	queryType, fieldBytes, err := lib.DecodeQueryType(queryDataBytes)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("failed to decode query data: %v", err))
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	// fetch data spec from store
	exists, err := k.Keeper.HasSpec(ctx, queryType)
	if !exists {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("can't decode query data for query type which doesn't exist in store: %v", queryType))
	}
	dataSpec, err := k.Keeper.GetSpec(ctx, queryType)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("failed to get data spec: %v", err))
	}
	// convert field bytes to string
	fields, err := lib.DecodeParamtypes(fieldBytes, dataSpec.QueryParameterTypes)

	return &types.QueryDecodeQuerydataResponse{Spec: fmt.Sprintf("%s: %s", queryType, fields)}, nil
}

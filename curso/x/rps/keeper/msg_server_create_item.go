package keeper

import (
	"context"

	"curso/x/rps/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateItem(goCtx context.Context, msg *types.MsgCreateItem) (*types.MsgCreateItemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	Item := types.Item{
		Name:        msg.Name,
		Description: msg.Description,
	}
	id, err := k.Keeper.CreateItem(ctx, Item)
	if err != nil {
		return nil, err
	}
	_ = id
	

	return &types.MsgCreateItemResponse{}, nil
}

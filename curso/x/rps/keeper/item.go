package keeper

import (
	"encoding/binary"
	"curso/x/rps/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


func GetItemIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) CreateItem(ctx sdk.Context, item types.Item) (uint64, error) {

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ItemKey))

	b := k.cdc.MustMarshal(&item)
	store.Set(GetItemIDBytes(1), b)
	return 1, nil
}

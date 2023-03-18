package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mkXultra/maki_chain/x/maki/types"
)

// GetMakiCount get the total number of maki
func (k Keeper) GetMakiCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MakiCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetMakiCount set the total number of maki
func (k Keeper) SetMakiCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MakiCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendMaki appends a maki in the store with a new id and update the count
func (k Keeper) AppendMaki(
	ctx sdk.Context,
	maki types.Maki,
) uint64 {
	// Create the maki
	count := k.GetMakiCount(ctx)

	// Set the ID of the appended value
	maki.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MakiKey))
	appendedValue := k.cdc.MustMarshal(&maki)
	store.Set(GetMakiIDBytes(maki.Id), appendedValue)

	// Update maki count
	k.SetMakiCount(ctx, count+1)

	return count
}

// SetMaki set a specific maki in the store
func (k Keeper) SetMaki(ctx sdk.Context, maki types.Maki) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MakiKey))
	b := k.cdc.MustMarshal(&maki)
	store.Set(GetMakiIDBytes(maki.Id), b)
}

// GetMaki returns a maki from its id
func (k Keeper) GetMaki(ctx sdk.Context, id uint64) (val types.Maki, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MakiKey))
	b := store.Get(GetMakiIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMaki removes a maki from the store
func (k Keeper) RemoveMaki(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MakiKey))
	store.Delete(GetMakiIDBytes(id))
}

// GetAllMaki returns all maki
func (k Keeper) GetAllMaki(ctx sdk.Context) (list []types.Maki) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MakiKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Maki
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetMakiIDBytes returns the byte representation of the ID
func GetMakiIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetMakiIDFromBytes returns ID in uint64 format from a byte array
func GetMakiIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

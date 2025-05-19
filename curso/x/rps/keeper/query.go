package keeper

import (
	"curso/x/rps/types"
)

var _ types.QueryServer = Keeper{}

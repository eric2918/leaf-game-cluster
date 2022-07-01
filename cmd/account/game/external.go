package game

import (
	"leaf-game-cluster/cmd/account/game/internal"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)

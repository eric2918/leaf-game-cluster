package game

import (
	"leaf-game-cluster/cmd/game/game/internal"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)

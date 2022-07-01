package game

import (
	"leaf-game-cluster/cmd/gate/game/internal"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)

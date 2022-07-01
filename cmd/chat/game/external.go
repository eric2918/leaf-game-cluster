package game

import (
	"leaf-game-cluster/cmd/chat/game/internal"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)

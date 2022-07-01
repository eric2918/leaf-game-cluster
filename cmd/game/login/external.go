package login

import (
	"leaf-game-cluster/cmd/game/login/internal"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)

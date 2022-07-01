package login

import (
	"leaf-game-cluster/cmd/account/login/internal"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)

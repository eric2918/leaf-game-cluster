package base

import (
	"leaf-game-cluster/conf"

	"github.com/eric2918/leaf-cluster/chanrpc"
	"github.com/eric2918/leaf-cluster/module"
)

func NewSkeleton() *module.Skeleton {
	skeleton := &module.Skeleton{
		GoLen:              conf.GoLen,
		TimerDispatcherLen: conf.TimerDispatcherLen,
		AsynCallLen:        conf.AsynCallLen,
		ChanRPCServer:      chanrpc.NewServer(conf.ChanRPCLen),
	}
	skeleton.Init()
	return skeleton
}

package main

import (
	"leaf-game-cluster/cmd/gate/game"
	"leaf-game-cluster/cmd/gate/gate"
	"leaf-game-cluster/cmd/gate/login"
	"leaf-game-cluster/conf"

	"github.com/eric2918/leaf-cluster"
	lconf "github.com/eric2918/leaf-cluster/conf"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ConsolePrompt = conf.Server.ConsolePrompt
	lconf.ProfilePath = conf.Server.ProfilePath
	lconf.ServerName = conf.Server.ServerName
	lconf.ListenAddr = conf.Server.ListenAddr
	lconf.ConnAddrs = conf.Server.ConnAddrs
	lconf.PendingWriteNum = conf.Server.PendingWriteNum
	lconf.HeartBeatInterval = conf.HeartBeatInterval

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}

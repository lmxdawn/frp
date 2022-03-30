package main

import (
	"github.com/fatedier/golib/crypto"
	"github.com/panglove/frp/pkg/auth"
	"github.com/panglove/frp/pkg/config"
	"github.com/panglove/frp/server"
	"testing"
	"time"
)

func Test2(t *testing.T) {
	dfg := config.GetDefaultServerConf()
	crypto.DefaultSalt = "frp"
	dfg.ServerConfig = auth.GetDefaultServerConf()
	dfg.DashboardAddr = "127.0.0.1"
	dfg.DashboardPort = 7500
	newFrpServer, _ := server.NewService(dfg)

	go newFrpServer.Run()

	for  {
		newFrpServer.GetProxyStatsByType("tcp")
		time.Sleep(time.Second*5)
	}
}

package main

import (
	"console"
	"elog"
	"netlib"
	//"os/signal"
	"runtime"
	"util"
)

const (
	CONN_NEW byte = iota
	// read write
	CONN_READ
	CONN_WRITE
	CONN_CLOSE
	EVENT_MAX
)

func main() {
	var (
		gameRun   = true
		gameFrame = 0
	)

	elog.InitLog(elog.INFO, true)

	initialize()
	listen := new(netlib.PeerListener)
	err := listen.Start("127.0.0.1:7702")
	util.CheckErrorCrash(err, "listen.Start")
	defer listen.Close()

	startup()
	elog.LogInfo("gameserver start...")
	for {
		if !gameRun {
			break
		}
		gameFrame++
		MainFrameRun(gameFrame)
	}
	elog.LogInfo("gameserver end...")
	elog.Flush()

}

func MainFrameRun(gameFrame int) {
	//log.Println(gameFrame)

}

func initialize() {
	elog.LogInfo("\nServer initializing...\n")

	// 开启多核
	runtime.GOMAXPROCS(runtime.NumCPU())
	elog.LogInfo("current CPUs: ", runtime.NumCPU(), "\n")

	// 开启控制台
	go console.Console()
	elog.LogInfo("Server console activated.\n")
}

func startup() {

}

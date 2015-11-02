package main

import (
	"fmt"
	"log"
	"net"
	"netlib"
	"util"
)

func main() {
	var (
		gameRun = true
	)

	initialize()

	listen := new(netlib.PeerListener)
	err := listen.Start("127.0.0.1:7700")
	util.CheckErrorCrash(err, "listen.Start")
	defer listen.Close()

	log.Println("gameserver start...")
	for {
		if !gameRun {
			break
		}
	}
	log.Println("gameserver end...")

}

func initialize() {
	fmt.Println("\nServer initializing...\n")

	// 开启多核
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Multi-CPU support active, current CPUs in use: ", runtime.NumCPU(), "\n")

	// 开启控制台
	go console.Console()
	fmt.Println("Server console activated.\n")

	// 准备通信线程
	go message.Message()
	fmt.Println("Message routines ready...\n")
}

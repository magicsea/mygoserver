package main

import (
	"fmt"
	"net"
	"util"
)

func main() {
	var (
		host   = "127.0.0.1"
		port   = "7000"
		remote = host + ":" + port
		//reader = bufio.NewReader(os.Stdin)
	)

	listen, err := net.Listen("tcp", remote)
	defer listen.Close()
	util.CheckErrorCrash(err)

	// 等待客户端连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("Accept error: ", err)
			continue
		}
		go connection.NewConnection(conn)
	}

	go onConnectEvent(listen)
	fmt.Println("gameserver start...")
}

func onConnectEvent(listen net.Listener) {
	fmt.Println("onconn")
}

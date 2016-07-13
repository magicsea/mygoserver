package main

import (
	"console"
	"elog"
	"fmt"
	"net"
	peer "netlib/peer"
	"serverApp"
)

type ConsoleCMD console.ConsoleCMD

//////////////////////////////////////////////////app
type simpleServer struct {
	serverApp.ServerAppBase
	test string
}

func NewServerApp(address string, loglevel int, cm []console.ConsoleCMD) *simpleServer {
	basePtr := serverApp.NewServerApp(address, loglevel, cm)
	app := simpleServer{*basePtr, "sss"}
	return &app
}

func (self *simpleServer) CreatePeer(conn net.Conn) *peer.PeerBase {
	return peer.NewPeer(conn)
}

/////////////////////////////////////////////////peer
type simplePeer struct {
	peer.PeerBase
}

func (p *simplePeer) OnRecvPack(code int, pack []byte) {
	elog.LogInfo("onRecv:")
}

//////////////////////////////main

func main() {
	app := NewServerApp("127.0.0.1:7700", elog.INFO, []console.ConsoleCMD{console.ConsoleCMD{"printme", printme}})
	app.Run()

}

func printme(a []string) {
	fmt.Print("printme=>" + a[0])
}

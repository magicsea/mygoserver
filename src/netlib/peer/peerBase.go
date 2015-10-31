package peer

import (
	"net"
)

type PeerBase struct {
	Id      int64
	RecvMsg chan string
	Socket  net.Conn
}

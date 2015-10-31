package netlib

import (
	"netlib/peer"
)

type PeerList []peer.PeerBase
type PeerListener struct {
	Address string
	Peers   PeerList
}

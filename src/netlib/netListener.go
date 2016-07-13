package netlib

import (
	"log"
	"net"
	"netlib/peer"
)

type PeerBase peer.PeerBase
type PeerList []peer.PeerBase
type PeerListener struct {
	Address        string
	IdCreator      int64
	Peers          PeerList
	Listen         net.Listener
	createPeerFunc CreatePeerCallBack
}

func (p *PeerListener) NewId() int64 {
	p.IdCreator++
	return p.IdCreator
}

type CreatePeerCallBack func(conn net.Conn) *peer.PeerBase

func (p *PeerListener) Start(address string, callback CreatePeerCallBack) error {
	p.Address = address
	p.createPeerFunc = callback
	//监听
	listen, err := net.Listen("tcp", p.Address)
	p.Listen = listen

	if err != nil {
		return err
	}

	// 等待客户端连接
	go p.StartAcceptPeer()
	return nil
}

func (p *PeerListener) Close() {
	p.Listen.Close()
}

//接收连接轮训
func (p *PeerListener) StartAcceptPeer() {
	for {
		conn, err := p.Listen.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
			continue
		}
		p.OnAcceptNewPeer(conn)
	}
}

//接收到新的连接
func (p *PeerListener) OnAcceptNewPeer(conn net.Conn) {
	log.Println("OnAcceptNewPeer", conn.RemoteAddr())

	newpeer := p.createPeerFunc(conn)
	newpeer.SetID(p.NewId())
	go newpeer.Run()
}

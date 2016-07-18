package peer

import (
	//"bufio"
	"io"
	"log"
	"net"
	pack "netlib/pack"
)

var (
	commingData = make([]byte, 1024)
)

type PeerBase struct {
	Pid       int64
	RecvMsg   chan NetPack
	writeMsg  chan NetPack
	Socket    net.Conn
	closeSign chan string
	readCache *pack.ReadyCache
}

func NewPeer(conn net.Conn) *PeerBase {
	msgChanR := make(chan NetPack, 10)
	msgChanW := make(chan NetPack, 10)
	closeSign := make(chan string)
	newpeer := &PeerBase{0, msgChanR, msgChanW, conn, closeSign, pack.NewReadyCache(msgChanR)}
	return newpeer
}

func (p *PeerBase) SetID(id int64) {
	p.Pid = id
}

//peer主线程
func (p *PeerBase) Run() {
	defer func() {
		conn.Close()
		if err := recover(); err != nil {
			printf("panic: %v\n\n%s", err, debug.Stack())
		}
	}()

	//rw work
	go p.ReadWork()
	go p.WriteWork()
	// peer main work
	for {
		select {
		case netpack := <-p.RecvMsg:
			p.OnRecvPack(netpack)
		case data := <-p.closeSign:
			p.OnDisconnect(data)
			goto EndPeer
		}
	}

EndPeer:
	p.OnDestory()
}

//读线程
func (p *PeerBase) ReadWork() {

	//buff := bufio.NewReader(p.Socket)

	for {
		lengh, err := p.Socket.Read(commingData)

		// 如果读入数据时出错，通知写回线程退出，广播退出信息，关闭用户连接
		if err != nil {
			var reason string
			if err == io.EOF {
				reason = " has left the game."
			} else {
				reason = "Read error: " + err.Error()
			}
			p.closeSign <- reason
			return
		}

		p.readCache.AddData(commingData[:lengh])
		//commingStr := string(commingData[:lengh])
		p.readCache.ParsePack()
		//p.RecvMsg <- commingStr

	}

}

//写线程
func (p *PeerBase) WriteWork() {
	for {
		msg := <-p.writeMsg
		p.Socket.Write(([]byte)(msg))
	}
}

func (p *PeerBase) OnRecvPack(netpack NetPack) {

}

func (p *PeerBase) Send(data string) {
	p.writeMsg <- data
}

func (p *PeerBase) SendPack(code int, pack []byte) {

}
func (p *PeerBase) FlushSend() {

}
func (p *PeerBase) OnDisconnect(reason string) {
	log.Println("Disconnect:", reason)
}

func (p *PeerBase) OnDestory() {
	log.Println("OnDestory", p.Pid)
}

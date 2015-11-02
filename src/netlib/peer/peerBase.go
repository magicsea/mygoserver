package peer

import (
	"io"
	"log"
	"net"
)

var (
	commingData = make([]byte, 1024)
)

type PeerBase struct {
	Pid     int64
	RecvMsg chan string
	Socket  net.Conn
}

func NewPeer(pid int64, conn net.Conn) *PeerBase {
	msgChan := make(chan string)
	newpeer := &PeerBase{pid, msgChan, conn}
	return newpeer
}

//开启消息读取线程
func (p *PeerBase) Run() {
	// 开始读入客户端信息
	for {
		lengh, err := p.Socket.Read(commingData)

		// 如果读入数据时出错，通知写回线程退出，广播退出信息，关闭用户连接
		if err != nil {
			if err == io.EOF {
				log.Println(" has left the game.", p.Pid)
			} else {
				log.Println("Read error: ", err)
			}
			p.Disconnect()
			return
		}

		commingStr := string(commingData[:lengh])
		log.Println("read:", commingStr)
	}
}

func (p *PeerBase) Disconnect() {
	log.Println("Disconnect:", p.Pid)
}

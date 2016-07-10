package serverApp

import (
	"console"
	"elog"
	"netlib"
	"runtime"
	"util"
)

type ServerAppBase struct {
	Address  string
	LogLevel int
	Quit     chan byte //结束标志
	cmds     []console.ConsoleCMD
}

func NewServerApp(address string, loglevel int, cmds []console.ConsoleCMD) *ServerAppBase {
	app := ServerAppBase{address, loglevel, make(chan byte), cmds}
	return &app
}

func (self *ServerAppBase) Run() {
	elog.InitLog(self.LogLevel, true)
	self.initialize()

	listen := new(netlib.PeerListener)
	err := listen.Start(self.Address)
	util.CheckErrorCrash(err, "listen.Start")
	defer listen.Close()

	self.onStart()
	<-self.Quit
	self.onEnd()
}

func (self *ServerAppBase) initialize() {
	elog.LogInfo("\nServer initializing...\n")

	// 开启多核
	runtime.GOMAXPROCS(runtime.NumCPU())
	elog.LogInfo("current CPUs: ", runtime.NumCPU(), "\n")

	// 开启控制台
	go console.Console(self.Quit, self.cmds)
	elog.LogInfo("Server console activated.\n")
}

func (self *ServerAppBase) onStart() {
	elog.LogInfo("gameserver start...")
}

func (self *ServerAppBase) onEnd() {
	elog.LogInfo("gameserver end...")
	elog.Flush()
}

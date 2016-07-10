package console

type CONSOLE_CALL func(args []string)

type ConsoleCMD struct {
	CMD      string
	Callback CONSOLE_CALL
}

type ConsoleCMDParse struct {
	cmdMap map[string]ConsoleCMD
}

func NewConsoleCMDParse(cmds []ConsoleCMD) *ConsoleCMDParse {
	mgr := ConsoleCMDParse{make(map[string]ConsoleCMD)}
	for i := 0; i < len(cmds); i++ {
		cmd := cmds[i]
		mgr.cmdMap[cmd.CMD] = cmd
	}
	return &mgr
}

func (self *ConsoleCMDParse) OnCMD(cmd string, arg []string) bool {
	cmdnew, ok := self.cmdMap[cmd]
	if ok {
		cmdnew.Callback(arg)
	}
	return ok
}

package console

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

var (
	cpus     = "cpus"
	routines = "gor"
	usernum  = "usernum"
	userlist = "userlist"
	memusage = "mem"
	quit     = "quit"
)

var (
	command = make([]byte, 1024)
	reader  = bufio.NewReader(os.Stdin)
)

func Console(quitFlag chan<- byte, cmds []ConsoleCMD) {
	var parse *ConsoleCMDParse
	if cmds != nil {
		parse = NewConsoleCMDParse(cmds)
	}
	for {
		command, _, _ = reader.ReadLine()
		strCmd := string(command)
		strs := strings.Split(strCmd, " ")
		if strs != nil && len(strs) >= 1 {
			cmd := strs[0]
			args := strs[1:]
			switch cmd {

			case quit:
				fmt.Println("Server stopped.")
				//os.Exit(0)
				quitFlag <- 1
			case cpus:
				fmt.Println("\nThe number of CPUs currently in use: ", runtime.NumCPU())

			case routines:
				fmt.Println("\nCurrent number of goroutines: ", runtime.NumGoroutine())

			case memusage:
				//us := &syscall.Rusage{}
				//err := syscall.Getrusage(syscall.RUSAGE_SELF, us)
				//if err != nil {
				//	fmt.Println("Get usage error: ", err, "\n")
				//} else {
				//	fmt.Printf("\nMemory Usage: %f MB\n\n", float64(us.Maxrss)/1024/1024)
				//}
			default:
				if parse != nil {
					if !parse.OnCMD(cmd, args) {
						fmt.Println("Command error, try again.")
					}
				}

			}
		}

	}
}

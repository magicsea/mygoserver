package util

import (
	"bufio"
	"elog"
	//"fmt"
	//"log"
	"os"
	//"runtime"
	//"runtime/debug"
)

func CheckErrorCrash(err error, addon string) {
	if err != nil {
		elog.LogError("fatal error:", addon, err.Error())

		//debug.PrintStack()
		//buf := make([]byte, 1<<20)
		//runtime.Stack(buf, true)
		//fmt.Printf("\n%s", buf)

		os.Exit(1)
	}
}

func Pause() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadByte()
}

func CheckError(err error) {
	if err != nil {
		elog.LogError("error:%s", err.Error())

		//debug.PrintStack()
	}
}

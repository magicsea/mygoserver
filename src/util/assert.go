package util

import (
	"fmt"
	"log"
	"os"
)

func CheckErrorCrash(err error, addon string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s.fatal error:%s", addon, err.Error())
		log.Fatal("fatal error:%s", err.Error())
		os.Exit(1)
	}
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error:%s", err.Error())
	}
}

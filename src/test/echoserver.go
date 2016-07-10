package main

import (
	"console"
	"elog"
	"fmt"
	"serverApp"
)

func main() {
	app := serverApp.NewServerApp("127.0.0.1:7700", elog.INFO, []console.ConsoleCMD{console.ConsoleCMD{"printme", printme}})
	app.Run()
}

func printme(a []string) {
	fmt.Print("printme=>" + a[0])
}

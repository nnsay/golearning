package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "mGo.log")
	file, err := os.OpenFile(LOGFILE, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	ilog := log.New(file, "iLog ", log.LstdFlags)
	fmt.Println("logfile:", LOGFILE)
	ilog.Println("hello world") // iLog 2022/01/07 00:30:23 hello world
	ilog.SetFlags(log.Lshortfile | log.LstdFlags)
	ilog.Println("你好") //iLog 2022/01/07 00:30:49 log.go:23: 你好
}

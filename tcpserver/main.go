package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var count = 0

func process(c net.Conn) {
	for {
		message, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println("ReadString error: ", err)
			break
		}

		if strings.TrimSpace(message) == "STOP" {
			fmt.Println("TCP client existing")
			break
		}

		fmt.Print("-> ", message)
		counter := "Client number: " + strconv.Itoa(count) + "\n"
		t := counter + " " + "Server Received At: " + time.Now().Format(time.RFC3339) + "\n"
		c.Write([]byte(t))
	}

	c.Close()
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("please provide host:port")
		return
	}

	connect := arguments[1]
	l, err := net.Listen("tcp", connect)
	if err != nil {
		fmt.Println("Listen error: ", err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Accept error: ", err)
			return
		}
		go process(c)
		count++
	}

	// for {
	// 	message, err := bufio.NewReader(c).ReadString('\n')
	// 	if err != nil {
	// 		fmt.Println("ReadString error: ", err)
	// 		return
	// 	}

	// 	if strings.TrimSpace(message) == "STOP" {
	// 		fmt.Println("TCP client existing")
	// 		return
	// 	}

	// 	fmt.Print("-> ", message)
	// 	t := "Server Received At: " + time.Now().Format(time.RFC3339) + "\n"
	// 	c.Write([]byte(t))
	// }

}

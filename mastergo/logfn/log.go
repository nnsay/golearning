package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("must give number")
		return
	}
	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	if number%2 == 0 {
		// log.Panic == log.Print + panic()
		// exit status 2
		log.Panicf("%d is ok", number)
	} else {
		// log.Fatal == log.Print + os.exit(1)
		// exit status 1
		log.Fatalf("%d is ok", number)
	}
}

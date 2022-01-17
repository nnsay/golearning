package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		log.Panic("Please give the server")
	}

	url := "https://" + arguments[1] + "/health"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	c := &http.Client{
		Timeout: 15 * time.Second,
	}
	res, err := c.Do(req)
	if err != nil {
		log.Panic(err)
	}
	if res == nil {
		log.Panic("res is nil")
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	log.Println("get result: ", string(data))
}

// go run restclient/main.go test2-api.neuralgalaxy.cn

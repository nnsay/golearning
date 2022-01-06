package main

import (
	"log"
	"log/syslog"
)

func main() {
	syslog, err := syslog.New(syslog.LOG_SYSLOG, "learn")
	if err != nil {
		log.Println("new syslog error", err)
		return
	}
	log.SetOutput(syslog)
	log.Println("system log record")
}

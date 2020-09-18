package golearning_test

import (
	"log"
	"testing"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)

	return func() {
		log.Printf("exit at %s (%s)", msg, time.Since(start))
	}
}

func TestExec(t *testing.T) {
	bigSlowOperation()
}

func print(t *testing.T, num int) {
	t.Log(num)
}

func TestMultipleDefer(t *testing.T) {
	defer print(t, 1)
	defer print(t, 2)
	defer print(t, 3)
}

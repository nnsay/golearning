package golearning_test

import (
	"fmt"
	"os"
	"runtime"
	"testing"
)

func f(x int) {
	// fmt.Printf("f(%d) \n", x+0/x)
	if x == 0 {
		panic(map[string]string{"errcode": "SV0001", "message": "x must great that 0"})
	}
	defer fmt.Printf("defer %d\n", x)
	// defer printStack()
	f(x - 1)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func TestMain(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			t.Logf("panic result: %#v", p)
		}
	}()

	f(3)
}

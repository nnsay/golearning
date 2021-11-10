package advancegrammar_test

import (
	"fmt"
	"testing"
	"time"
)

// CSP: communicating sequential processes
// csp vs actor
// 1. csp通过channel进行通讯, actor直接通讯, csp更松耦合一些
// 2. channel有容量限制且独立于groutine, 而actor的mailbox容量无限, 接收进程总是被动的处理消息

func service1() string {
	time.Sleep(time.Millisecond * 500)
	return "sevice1 done"
}

func service2() {
	fmt.Println("working on the service2")
	time.Sleep(time.Millisecond * 500)
	fmt.Println("server2 done")
}

// 默认串行
func TestService(t *testing.T) {
	fmt.Println(service1())
	service2()
}

func asyncService1() chan string {
	// 声明 chan/map/set
	ch := make(chan string)

	go func() {
		ret := service1()
		fmt.Println("return the service1")
		ch <- ret
		fmt.Println("service1 exited")
	}()

	return ch
}

func TestAsyncService1(t *testing.T) {
	ch := asyncService1()
	service2()
	fmt.Println(<-ch)
}

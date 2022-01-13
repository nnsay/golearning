package advancegrammar_test

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "sevice done"
}

func asyncService() chan string {
	// 声明 chan/map/set
	ch := make(chan string)

	go func() {
		ret := service()
		fmt.Println("return the service")
		ch <- ret
		fmt.Println("service exited")
	}()

	return ch
}

// select控制超时
func TestSelect(t *testing.T) {
	select {
	case result := <-asyncService():
		t.Logf("success result: %s\n", result)
	case <-time.After(time.Millisecond * 400):
		t.Errorf("timeout error\n")
		// default:
		// 	t.Logf("default error")
	}
}

// https://gobyexample.com/non-blocking-channel-operations
// 这个文章说select是阻塞, 但是通过加default可以变成非阻塞, 刚才也有同学说close通道后<-chan就是非阻塞了, 所以大概理论也许是:
// 1. 不关闭chan时, select在其阻塞逻辑块中得知通道正常, 只是没有可用的消息, 所选择走default, 所以通道标记为没有关闭;
// 2. 关闭chan时, select在其阻塞逻辑块中得知通道关闭,<-chan非阻塞, 此时case直接优先命中, 所以走了case的逻辑, 即此时通道标记关闭
func TestSelectCancel(t *testing.T) {
	ch := make(chan int)
	close(ch)
	select {
	case <-ch:
		t.Logf("close")
	default:
		t.Logf("no close")
	}
	t.Logf("done")
}

// 3. select是阻塞块, select块没有完成其下逻辑也就不能完成
func TestSelectBlock(t *testing.T) {
	ch := make(chan int, 1)
	t.Logf("1\n")
	ch <- 1
	t.Logf("2\n")
	// close(ch)
	select {
	case <-ch:
		t.Logf("done\n")
	case <-time.After(time.Second * 1):
		t.Logf("timeout\n")
		// default:
		// 	t.Logf("no close")
	}
	t.Logf("done\n")
}

// 1. push chan是一个阻塞操作
// 2. pop chan是一个阻塞操作
// 3. 以 go 声明的协程方法是非阻塞编程
func TestNilChan(t *testing.T) {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	t.Log("make chan done")
	result := <-c
	t.Log("wait the chan value")
	t.Logf("result: %d", result)
}

func abortLaunch() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
}

func launch() {
	fmt.Printf("Commencing coundown\n")
	tick := time.Tick(1 * time.Second)
	for counterdown := 10; counterdown > 0; counterdown-- {
		fmt.Printf("%d\n", counterdown)
		<-tick
	}
}
func TestLaunch(t *testing.T) {
	launch()
	t.Logf("launch...\n")
}

func TestAbortLaunch(t *testing.T) {
	fmt.Printf("Commencing coundown\n")
	tick := time.NewTicker(1 * time.Second)
	abort := make(chan struct{})

	// 使用协程监控标准输入
	// go func() {
	// 	os.Stdin.Read(make([]byte, 1))
	// 	abort <- struct{}{}
	// }()

	for counterdown := 10; counterdown > 0; counterdown-- {
		fmt.Printf("%d\n", counterdown)
		select {
		case <-tick.C:
			// return
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	tick.Stop()
	t.Logf("launch...\n")
}

func TestNilChan2(t *testing.T) {
	var wg sync.WaitGroup

	add := func(c chan int) {
		sum := 0
		t := time.NewTimer(time.Second)

		for {
			select {
			case input := <-c:
				sum = sum + input
			case <-t.C:
				c = nil
				fmt.Println(sum)
				wg.Done()
			}
		}
	}
	send := func(c chan int) {
		for {
			c <- rand.Intn(10)
		}
	}

	c := make(chan int)
	rand.Seed(time.Now().Unix())
	wg.Add(1)
	go add(c)
	go send(c)
	wg.Wait()
}

package advancegrammar_test

import (
	"runtime"
	"testing"
	"time"
)

func runTask() time.Time {
	time.Sleep(time.Millisecond * 10)
	return time.Now()
}

// 1. chan不设置长度时会阻塞: 放消息阻塞, 取消息阻塞, 也就是Buffer长度是1
// 2. 当第一消息收到后立刻退出程序
// 3. 问题: 有大量其他协程任然在内存中, 容易造成协程泄露
// 4. 使用有缓存的chan
func doRace() int {
	currencyNum := 10
	ch := make(chan int)
	for i := 0; i < currencyNum; i++ {
		go func(idx int) {
			runTask()
			ch <- idx
		}(i)
	}
	return 1
}
func TestRace(t *testing.T) {
	t.Logf("before goroutine count: %d", runtime.NumGoroutine())
	t.Logf("result: %d", doRace())
	time.Sleep(time.Second * 2)
	t.Logf("after goroutine count: %d", runtime.NumGoroutine())
}

func TestChanPopBlock(t *testing.T) {
	t.Log("1")
	ch := make(chan int)
	t.Log("2")
	<-ch
	t.Log("done")
}

func TestChanPushBlock(t *testing.T) {
	t.Log("1")
	ch := make(chan int)
	t.Log("2")
	ch <- 1
	t.Log("3")
	ch <- 2
	t.Log("done")
}

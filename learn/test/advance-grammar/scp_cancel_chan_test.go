package advancegrammar_test

import (
	"fmt"
	"sync"
	"testing"
)

func isCancel(cancelCh chan struct{}) bool {
	select {
	case <-cancelCh:
		return true
	default:
		return false
	}
}

// func cancel1(cancelCh chan struct{}) {
// 	cancelCh <- struct{}{}
// }

func cancel2(cancelCh chan struct{}) {
	close(cancelCh)
}

func TestCancel(t *testing.T) {
	var mu sync.WaitGroup
	cancelCh := make(chan struct{})

	for i := 0; i < 5; i++ {
		mu.Add(1)
		go func(ii int, cc chan struct{}, wg *sync.WaitGroup) {
			for {
				fmt.Printf("goroutine %d is running\n", ii)
				if isCancel(cc) {
					break
				}
			}
			fmt.Printf("goroutine %d is canceled\n", ii)
			wg.Done()
		}(i, cancelCh, &mu)
	}

	// 通过消息传递取消因为只有一个消息, 则也只有一个携程会被取消
	// cancel1(cancelCh)
	// 通过关闭通道, 可以让每个协程都知道通道关闭, 从而达到所有协程取消的目的
	// 通道关闭时, 通道继续<-, 则返回类型的零值

	// mu.Add(1)
	go func(ch chan struct{}, wg *sync.WaitGroup) {
		cancel2(ch)
		fmt.Println("closed")
		// wg.Done()
	}(cancelCh, &mu)

	mu.Wait()
	fmt.Println("all done")
}

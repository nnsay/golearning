package advancegrammar_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 知识
// 1. 往关闭的chan中推数据会引发panic
// 2. 从chan消费数据时有两个参数: 值,通道是否开通
// 3. 通道关闭时继续取数据, 则立刻返回通道类型的零值
// 4. 通道关闭的消息针对所有引用通道的地方广播

// 生产者主动关闭
func dataProducer(ch chan int, mu *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		mu.Done()
	}()
}

// 消费者判断是否关闭
func dataConsumer(name string, ch chan int, mu *sync.WaitGroup) {
	go func(tag string) {
		for {
			data, ok := <-ch
			if !ok {
				break
			}
			fmt.Printf("[%s]comsumer: %d\n", tag, data)
		}
		mu.Done()
	}(name)
}

func TestCloseChan(t *testing.T) {
	ch := make(chan int)
	var mu sync.WaitGroup

	// 生产者:消费者=> 1:N
	mu.Add(1)
	dataProducer(ch, &mu)
	mu.Add(1)
	dataConsumer("A", ch, &mu)
	mu.Add(1)
	dataConsumer("B", ch, &mu)
	mu.Add(1)
	dataConsumer("C", ch, &mu)
	mu.Wait()
}

func TestCloseChanBradcase(t *testing.T) {
	ch := make(chan int)
	go func(c chan int) {
		c <- 1
		close(c)
	}(ch)

	time.Sleep(time.Millisecond * 500)

	go func(c chan int) {
		data, ok := <-c
		t.Log("1", ok, data)
	}(ch)
	go func(c chan int) {
		data, ok := <-c
		t.Log("2", ok, data)
	}(ch)
	go func(c chan int) {
		data, ok := <-c
		t.Log("3", ok, data)
	}(ch)
	time.Sleep(time.Millisecond * 500)
}

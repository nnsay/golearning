package golearning_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 协程并发
// 1. 方法前不加go则是普通方法, 顺序打印
// 2. 加上go就是将函数放入协程执行, 乱序打印
// 3. 协程函数注意闭包
func TestGoroutine(t *testing.T) {
	for i := 1; i <= 10; i++ {
		go func(i int) {
			fmt.Println(i * 1)
		}(i)
	}
	time.Sleep(time.Microsecond * 50)
}

// 并发问题: 线程不安全
// 协程对于counter的自增不是一个线程安全的操作, 所以5000次执行后结果不一定是5000
func TestGoroutineLock(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second * 2)
	t.Logf("counter=%d", counter)
}

// 线程安全解决方案: 并发锁
func TestGoroutineThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 1; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()

			mut.Lock()
			counter++
		}()
	}
	time.Sleep(time.Second * 2)
	t.Logf("counter=%d", counter)
}

// 并发组等待
func TestGoroutineWaitGroup(t *testing.T) {
	// 好奇: 1. 声明后直接使用变量, 貌似没有赋值或者初始化, 但是go有零值初始化, 所以这里申明了就可以直接使用
	var wg sync.WaitGroup // 零值对象
	var mut sync.Mutex    // 零值对象
	var counter int       // 零值: 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(index int) {
			defer func() {
				mut.Unlock()
			}()
			t.Logf("%d [r]current value %d", index, counter)
			mut.Lock()
			t.Logf("%d [w]current value %d", index, counter)
			counter++
			wg.Done()
		}(i)
	}
	wg.Wait()
	t.Logf("counter=%d", counter)
}

// mutex是读写互斥的锁, 大部分情况我们写互斥读不互斥, 改写mutex
func TestGoroutineRWLock(t *testing.T) {
	var wg sync.WaitGroup // 零值对象
	var mut sync.RWMutex  // 零值对象
	var counter int       // 零值: 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(index int) {
			defer func() {
				mut.Unlock()
			}()
			t.Logf("%d [r]current value %d", index, counter)
			mut.Lock()
			t.Logf("%d [w]current value %d", index, counter)
			counter++
			wg.Done()
		}(i)
	}
	wg.Wait()
	t.Logf("counter=%d", counter)
}

package advancegrammar_test

import (
	"sync"
	"testing"
	"unsafe"
)

// 单例模式
type Signleton struct {
	name string
	age  int
}

// 以下三种声明once方式结果一致
// var once = *new(sync.Once)
// var once = new(sync.Once)
var once sync.Once

var singleton *Signleton

func getSingleton() *Signleton {
	once.Do(func() {
		singleton = new(Signleton)
	})
	return singleton
}

func TestSingleton(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			s := getSingleton()
			t.Logf("address: %d", unsafe.Pointer(s))
			wg.Done()
		}()
	}
	wg.Wait()
}

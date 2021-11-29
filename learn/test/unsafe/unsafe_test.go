package unsafe

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"unsafe"
)

type MyInt int

// 不安全的类型转化
func TestTypeConvert(t *testing.T) {
	// 成功转化的
	num := []int{1, 2, 3, 4, 5}
	myNum := *(*[]MyInt)(unsafe.Pointer(&num))
	intNum := *(*[]int)(unsafe.Pointer(&num))
	t.Logf("myNum: %#v", myNum)
	t.Logf("intNum: %#v", intNum)

	// 错误转化的
	// i := 10
	// f := *(*[]float32)(unsafe.Pointer(&i))
	// t.Logf("f: %#v", f)
}

// 原子类型操作
func TestAtomic(t *testing.T) {
	var shareBuffer unsafe.Pointer
	writeData := func() {
		data := []int{}
		for i := 0; i < 10; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&shareBuffer, unsafe.Pointer(&data))
	}
	readData := func() {
		data := atomic.LoadPointer(&shareBuffer)
		fmt.Println(data, *(*[]int)(data))
	}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			writeData()
			wg.Done()
		}()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			readData()
			wg.Done()
		}()
	}

	wg.Wait()
}

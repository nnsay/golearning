package golearning_test

import (
	"fmt"
	"testing"
	"time"
)

// 定义类型
type TestFn = func(op ...int) int

func logWrapper(fn TestFn) TestFn {
	return func(op ...int) int {
		start := time.Now()
		rst := fn(op...) // 解析数组
		fmt.Println("run time: ", time.Since(start).Seconds())
		return rst
	}
}

// 可变参数
func slowFn(op ...int) int {
	time.Sleep(time.Second * 2)
	rst := 0
	for _, v := range op {
		rst += v
	}
	return rst
}
func TestWrapperLog(t *testing.T) {
	logFn := logWrapper(slowFn)
	rst := logFn(10, 11, 21, 43)
	t.Logf("log fn result: %d", rst)
}

func TestDefer(t *testing.T) {
	// defer类似于try-catch-finally中的finally
	defer func() {
		t.Logf("come from defer")
	}()

	t.Logf("start")
	panic("errr")
}

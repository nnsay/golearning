package golearning_test

import (
	"errors"
	"fmt"
	"testing"
)

// os.exit
// 1. 不会执行defer
// 2. 不会打印堆栈信息
// panic
// 1. 会执行defer
// 2. 如果有defer执行defer, 否则打印调用栈
// recover
// 1. recover后进行有效的恢复操作, 不仅仅是打印日志, 否则容易制造僵尸进程
func TestPanicVsExit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from", err)
		}
	}()
	fmt.Println("Start")
	// os.Exit(-1)
	panic(errors.New("some thing was wrong"))
	// fmt.Println("End")
}

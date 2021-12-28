package advancegrammar_test

import (
	"fmt"
	"testing"
)

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
	fmt.Println("counter close")
}

// 单向chan: chan<- 只写; <-chan只读;
// 双向chan可以转化为单向chan(隐式转化), 但是单向chan不能转化为双向
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
	fmt.Println("squarer close")
}

// 使用range自动再chan关闭且没有值时退出
func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func TestMain(t *testing.T) {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

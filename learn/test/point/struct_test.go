package golearning_test

import (
	"fmt"
	"testing"
)

type composer struct {
	name      string
	birthYear int
}

func Test(t *testing.T) {
	t.Log("hlelo")
}

func TestCreate(t *testing.T) {
	c1 := composer{name: "c1name", birthYear: 10}
	c11 := composer{"c1name", 10}
	c2 := new(composer)
	c2.name = "c2name"
	c2.birthYear = 11
	c3 := &composer{name: "c3name", birthYear: 12}
	t.Log(c1, c11)
	t.Log(c2, c3)
}

func inflate(s []int, fact int) {
	for i, v := range s {
		fmt.Printf("i=%d, v=%.2d\n", i, v)
		s[i] *= fact
	}
}
func TestRangeSlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 11, 12}
	t.Log(s)
	inflate(s, 3)
	t.Log(s)
}

func TestArrayAndSet(t *testing.T) {
	arr := new([7]string)
	t.Logf("arr type: %T", arr)
	s1 := arr[:]
	t.Logf("set type: %T", s1)
}

type Product struct {
	name  string
	price float64
}

func (product Product) String() string {
	return fmt.Sprintf("%s(%.2f)", product.name, product.price)
}

func TestComplexSet(t *testing.T) {
	// 结构体指针类型初始化: 没有使用&, go可以自动识别
	products := []*Product{{"apple", 19.4}, {"orige", 0.27}, {"banana", 99}}
	// products := []*Product{&Product{"apple", 19.4}, &Product{"orige", 0.27}, &Product{"banana", 99}}
	t.Log(products, "\n")
	for _, product := range products {
		product.price += 0.50
	}
	t.Log(products, "\n")
}

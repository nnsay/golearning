package golearning_test

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	// 使用 make
	var runes []rune
	t.Logf("emtpy: %v", runes)
	s1 := make([]int, 4, 4)
	t.Logf("the s1: %v", s1)
	s1 = append(s1, 4)
	t.Log(s1, len(s1), cap(s1))

	// 使用类型
	var s2 []string
	t.Log(s2, len(s2), cap(s2))

	// 使用类型
	s3 := []string{"Hello"}
	t.Log(s3, len(s3), cap(s3))
}

func TestArraySlice(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	// 使用截取器, 从数组创建切片
	arr2 := arr[:]
	arr2 = append(arr2, 6)
	t.Log(arr2, len(arr2), cap(arr2))
	t.Log(arr2[1:])
	t.Log(arr2[:len(arr2)-2])

	// 遍历
	for idx, item := range arr2 {
		t.Logf("index: %v, value: %v", idx, item)
	}
}

// 切片是如何扩容的: 512之前都是2倍扩容, 过了512后扩容变缓(四分之一左右)
func TestHowToCap(t *testing.T) {
	arr := make([]int, 0)
	for i := 0; i < 2000; i++ {
		t.Logf("len: %v, cap %v \n", len(arr), cap(arr))
		arr = append(arr, i)
	}
}

// 切片是引用传递(实际是值传递, 传递的是header)
// 简单地说： slice 作为参数传递给函数其实是传递 slice 的值，这个值被称作一个 header ，它只包含了一个指向底层数组的指针。
// 当向函数传递一个 slice ，将复制一个 header 的副本，这个副本包含一个指向同一个底层数组的指针。
// 修改 slice 的元素间接地修改底层数组的元素，也就是所有指向同一个底层数组的 slice 会响应这个变化，主函数的 slice 也就一同修改了 s[0] 的值。
func TestSliceValuePass(t *testing.T) {
	arr := []int{5, 6}
	aa := changeSlice(arr) // 引用传递, 会影响原始数组
	t.Logf("origin slice: %v\n", arr)
	t.Logf("change slice: %v\n", aa)

	a := [5]int{1, 2, 3, 4, 5}
	s := a[:]

	fmt.Printf("%p \n", &s)    // 0xc000118240
	fmt.Printf("%p \n", &s[0]) // 0xc000132240
	modify2(s)
	fmt.Println(s[0])
}

func changeSlice(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}
	return arr
}

func modify2(s []int) {
	fmt.Printf("%p \n", &s)    // 0xc000118258
	fmt.Printf("%p \n", &s[0]) // 0xc000132240
	s[0] += 100
	fmt.Printf("%p \n", &s)    // 0xc000118258
	fmt.Printf("%p \n", &s[0]) // 0xc000132240
}

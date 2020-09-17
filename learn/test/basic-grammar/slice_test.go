package golearning_test

import "testing"

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

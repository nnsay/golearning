package suanfa

import (
	"fmt"
	"testing"
)

// 二分查找: 每次取中间值和目标数字比较, 如果相等则位置在中间值, 否则继续向前/后取中间值, 如果边界重合还未找到则不存在与数组
// 注意: mid=(high+low)/2

func binarySearch(arr []int, num int) int {
	idx := -1
	low := 0
	high := len(arr) - 1

	for {
		midIndex := (high + low) / 2
		midValue := arr[midIndex]
		if low >= high {
			break
		}
		fmt.Printf("current mid %d\n", midIndex)
		if num == midValue {
			idx = midIndex
			break
		} else if num < midValue {
			high = midIndex - 1
		} else if num > midValue {
			low = midIndex + 1
		}
	}

	return idx
}

func TestVerify(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	num := -1
	idx := binarySearch(arr, num)
	t.Logf("%d at %d", num, idx)
}

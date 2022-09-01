package suanfa

import (
	"fmt"
	"testing"
)

// 选择排序: 遍历N次数组, 每一次遍历数组找到最大的元素并将其移动到末端
// 0, 0, n-1
// 1, 0, n-2
// 2, 0, n-3
// i, 0, n-i-1
func selectSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	for i := 0; i < len(arr); i++ {
		maxValIdx := 0
		for j := 0; j < len(arr)-i; j++ {
			if j != maxValIdx && arr[j] > arr[maxValIdx] {
				maxValIdx = j
			}
		}
		fmt.Printf("%d: arr[%d]=%d\n", i, maxValIdx, arr[maxValIdx])
		arr[len(arr)-i-1], arr[maxValIdx] = arr[maxValIdx], arr[len(arr)-i-1]
	}
}

// 冒泡排序: 相邻两个元素比较, 小的左移大的右移, 每进行一次则把元素向左右两边归纳一次,
// 因为每一次遍历至少有一个元素找到正确位置, 所以移动 N 次能保证所有元素有序
func bubbleSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

// 插入排序: 每新来一个元素从现有数组右侧插入, 如果新元素小于现有数组右侧, 则左移动直到找到一个合适的位置, 如扑克牌整理
func insertSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	for i := 1; i < len(arr); i++ { // 新来的元素
		if arr[i] < arr[i-1] {
			for cur := i; cur > 0; cur-- { // 现有数组变量
				if arr[cur] < arr[cur-1] {
					arr[cur-1], arr[cur] = arr[cur], arr[cur-1]
				}
			}
		}
	}
}

func TestSort(t *testing.T) {
	arr := []int{3, 1, 4, 7, 9, 6, -1, 2, 7, -10}

	// arr := []int{3, 7, 8, 5, 9}
	t.Logf("sort before: %#v", arr)
	// selectSort(arr)
	// bubbleSort(arr)
	insertSort(arr)
	t.Logf("sort after : %#v", arr)
}

package shuxue

import (
	"fmt"
	"testing"
)

// 10进制转二进制
func TestDo102(t *testing.T) {
	results := do102(12)
	t.Logf("二进制结果: %v", results)
}

func do102(n int) []int {
	var results = make([]int, 0)
	var lefts = n % 2
	var times = n / 2
	results = append(results, lefts)
	fmt.Printf("times: %d, lefts: %d\n", times, lefts)
	for times != 0 {
		lefts = times % 2
		times = times / 2
		results = append(results, lefts)
		fmt.Printf("times: %d, lefts: %d\n", times, lefts)
	}
	return results
}

package golearning_test

import (
	"fmt"
	"testing"
)

func GenIncreFn() func(int) int {
	var index int

	return func(by int) int {
		index += by
		return index
	}
}

func TestClosures(t *testing.T) {
	increFn := GenIncreFn()
	increFn(1)
	increFn(2)
	result := increFn(3)
	t.Logf("result: %v", result)
}

func Sum(values ...int) int {
	fmt.Printf("values type: %T", values)
	var total int
	for _, value := range values {
		total += value
	}
	return total
}

func TestMultipleParameters(t *testing.T) {
	result := Sum(1, 2, 3, 4, 5)
	t.Logf("all parameter total: %d", result)
}

package golearning_test

import (
	"errors"
	"testing"
)

var ErrorLessThan2 = errors.New("n should >2")
var ErrorLargerThan100 = errors.New("n should <100")

func GetFibonacci(n int) ([]int, error) {
	if n < 0 {
		return nil, ErrorLessThan2
	}
	if n > 100 {
		return nil, ErrorLargerThan100
	}
	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}

	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	v, err := GetFibonacci(-10)
	// 尽早失败
	if err != nil {
		t.Errorf("error: %v", err)
		return
	}
	t.Logf("result: %v", v)
}

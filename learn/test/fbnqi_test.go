package arrayt

import (
	"testing"
)

func TestFb(t *testing.T) {
	num := 3
	t.Logf("result: %d", getFb(num))
}

func getFb(num int) (result int) {
	first := 0
	second := 1
	if num <= 0 {
		result = 0
		return
	}
	if num == 1 {
		result = first
		return
	}
	if num == 2 {
		result = second
		return
	}
	result = getFb(num-2) + getFb(num-1)
	return
}

package golearning_test

import "testing"

func TestGCD(t *testing.T) {
	rst := gcd(32, 4)
	t.Log(rst)
	rst = fcb(3)
	t.Log(rst)

	rst = fcb2(3, 3)
	t.Log(rst)
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fcb(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func fcb2(i int, n int) int {
	i *= 2
	n--
	if n != 0 {
		return fcb2(i, n)
	}
	return i
}

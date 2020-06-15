package gocn

import (
	"testing"
)

func TestNil(t *testing.T) {
	x := interface{}(nil)
	y := (*int)(nil)
	a := y == x
	b := y == nil
	_, c := x.(interface{})
	t.Log(a, b, c) // false true false
}

func TestProgram(t *testing.T) {
	var s []int
	t.Logf("empty slice length: %v, and the caption is: %v", len(s), cap(s))
	s = append(s, 1)
	t.Log(s)

	var m map[string]int
	// m = make(map[string]int)
	t.Logf("empty slice length: %v", len(m))

	m["one"] = 1 // assignment to entry in nil map
	t.Log(m)
}

func TestMap(t *testing.T) {
	a := make([]map[string]int, 100)
	for i := 0; i < 100; i++ {
		a[i] = make(map[string]int)
		a[i]["id"] = i // 假如map没有被初始化,则这里报错: assignment to entry in nil map
		a[i]["investor"] = i
		// a[i] = map[string]int{"id": i, "investor": i}
	}
	t.Log(a)
}

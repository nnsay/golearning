package golearning_test

import (
	"fmt"
	"testing"
)

type composer struct {
	name      string
	birthYear int
}

func TestCreate(t *testing.T) {
	c1 := composer{name: "c1name", birthYear: 10}
	c11 := composer{"c1name", 10}
	c2 := new(composer)
	c2.name = "c2name"
	c2.birthYear = 11
	c3 := &composer{name: "c3name", birthYear: 12}
	t.Log(c1, c11)
	t.Log(c2, c3)
}

func inflate(s []int, fact int) {
	for i := range s {
		fmt.Printf("i=%d\n", i)
		s[i] *= fact
	}
}
func TestRangeSlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	t.Log(s)
	inflate(s, 3)
	t.Log(s)
}

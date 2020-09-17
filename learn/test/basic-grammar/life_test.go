package golearning_test

import "testing"

func TestLoopVar(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Logf("the value is %v", i)
	}
	// t.Logf("after loop the value is: %v", i) // undefined: i
}

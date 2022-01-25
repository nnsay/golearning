package golearning_test

import (
	"runtime"
	"testing"
)

func TestRuntime(t *testing.T) {
	t.Logf("1: %s \n", runtime.GOOS)
	t.Logf("2: %s \n", runtime.GOARCH)
	t.Logf("2: %s \n", runtime.Version())
}

package golearning_test

import (
	"path/filepath"
	"testing"
)

func TestFileName(t *testing.T) {
	file := "/Users/wangjian/Downloads/debug/data server zip/rest(tr3-start0)/subject1_pre_45frames_withrecon.zip"
	dir := filepath.Base(file)
	t.Logf("1: %s \n", dir)
}

package golearning_test

import (
	"fmt"
	"os/exec"
	"syscall"
	"testing"
)

func TestClosures(t *testing.T) {
	cmd := exec.Command("/Users/wangjian/Workspace/practice/golearning/learn/test/dockerauth/extensions/ext_auth.sh")
	output, err := cmd.Output()
	es := 0
	et := ""
	if err == nil {
	} else if ee, ok := err.(*exec.ExitError); ok {
		es = ee.Sys().(syscall.WaitStatus).ExitStatus()
		et = string(ee.Stderr)
	} else {
		es = 100
		et = fmt.Sprintf("cmd run error: %s", err)
	}
	t.Logf("output: %#v, lenght: %#v, es: %#v, et: %#v", output, len(output), es, et)
}

package unittest

import "testing"

// 1. go test -v -cover 使用cover提供函数测试覆盖率
// 2. 断言
func TestFunction(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expects := [...]int{1, 4, 9}
	t.Log(1)
	for i := 0; i < len(inputs); i++ {
		v := Square(inputs[i])
		if v != expects[i] {
			t.Errorf("acture value: %d, expect value: %d", v, expects[i])
		}
	}
}

// Error 不产生中断
func TestError(t *testing.T) {
	t.Log("start")
	t.Error("error")
	t.Logf("end")
}

// Fail 不产生中断
func TestFail(t *testing.T) {
	t.Log("start")
	t.Fail()
	t.Logf("end")
}

// FailNow 产生中断
func TestFailNow(t *testing.T) {
	t.Log("start")
	t.FailNow()
	t.Logf("end")
}

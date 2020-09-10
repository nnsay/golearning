package arrayt

import "testing"

func TestPoint(t *testing.T) {
	p := new(int)
	t.Logf("the value point: %v", p)
	t.Logf("the zore value: %v", *p)
	*p = 100
	t.Logf("the value variable: %v", *p)
}

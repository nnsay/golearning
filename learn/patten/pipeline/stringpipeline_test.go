package pipeline

import "testing"

func TestStringPipeline(t *testing.T) {
	covFilter := &CovFilter{}
	splitFilter := &SplitFilter{}
	sumFilter := &SumFilter{}
	sp := NewStraightPipeline("MyTest", splitFilter, covFilter, sumFilter)
	ret, err := sp.Process("1,2,3,4")
	if err != nil {
		t.Error(err)
	}
	t.Logf("result: %v", ret)
}

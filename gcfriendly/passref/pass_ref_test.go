package passref

import "testing"

const NumberElems = 1000

type Content struct {
	Detail [10000]int
}

func withValue(arr [NumberElems]Content) int {
	return 0
}

func withReference(arr *[NumberElems]Content) int {
	return 0
}

func TestFn(t *testing.T) {
	var arr [NumberElems]Content

	withValue(arr)
	withReference(&arr)
}

func BenchmarkWithValue(b *testing.B) {
	var arr [NumberElems]Content
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withValue(arr)
	}
	b.StopTimer()
}

func BenchmarkWithReference(b *testing.B) {
	var arr [NumberElems]Content
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withReference(&arr)
	}
	b.StopTimer()
}

/*
	1. GC打印日志: GODEBUG=gctrace=1
	GODEBUG=gctrace=1 go test -bench=.
	2. trace 工具
	go test -bench=BenchmarkWithReference -trace=ref.out
	go test -bench=BenchmarkWithValue -trace=val.out
	go tool trace val.out
	go tool trace ref.out
**/

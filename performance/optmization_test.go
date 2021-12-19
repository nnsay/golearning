package performance

import "testing"

func TestProcessRequest(t *testing.T) {
	reqs := []string{}
	reqs = append(reqs, createRequest())
	reps := processRequest(reqs)
	t.Logf(reps[0])
}

// go test -bench=. -cpuprofile=cpuprofile
// go test -bench=. -memprofile=memprofile
// easyjson -all structs.go
// go tool pprof cpuprofile
// 	top
// 	top -cum
//  list processRequest
func BenchmarkProcessRequest(b *testing.B) {
	reqs := []string{}
	reqs = append(reqs, createRequest())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		processRequest(reqs)
	}
	b.StopTimer()
}

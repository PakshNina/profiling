package factorial

import "testing"

const (
	factNumber = 3628800
)

// go test -v ./... -bench=. -benchtime 100x (-count=2)
func BenchmarkRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Recursive(factNumber)
	}
}

func BenchmarkDynamic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dynamic(factNumber)
	}
}

// go test ./... -bench="BenchmarkCalculate" -run=^# -count=2 | tee old.txt
// benchstat old.txt new.txt
// go test -bench='BenchmarkCalculate' -cpuprofile='cpu.prof' -memprofile='mem.prof'
// go tool pprof cpu.prof
func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate()
	}
}

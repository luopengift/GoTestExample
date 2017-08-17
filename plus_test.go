package GoTestExample

import (
	"testing"
)

func Benchmark_plusEqual(b *testing.B) {
    cnt := 0
	for i := 0; i < b.N; i++ {
		cnt += 1
	}
}
func Benchmark_plusOne(b *testing.B) {
    cnt := 0
	for i := 0; i < b.N; i++ {
	    cnt = cnt + 1
    }
}

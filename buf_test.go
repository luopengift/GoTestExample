package GoTestExample

import (
	"testing"
)

const (
	B = 8
	K = 1024 * B
	M = 1024 * K
	G = 1024 * M

	MAX = 1 * K
)

func Benchmark_Buffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewBuffer(MAX)
	}
}
func Benchmark_Byte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewByte(MAX)
	}
}

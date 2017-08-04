package GoTestExample

import (
	"testing"
        "bytes"
)

const (
	B = 8
	K = 1024 * B
	M = 1024 * K
	G = 1024 * M

	MAX = 1 * K
)

func NewBuffer(max int) *bytes.Buffer {
        return bytes.NewBuffer(make([]byte, 0, max))
}

func NewByte(max int) []byte {
        return make([]byte, 0, max)
}

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

package GoTestExample

import (
	"bytes"
	"testing"
    "github.com/luopengift/types"
)

func NewBuffer(max int) *bytes.Buffer {
	return bytes.NewBuffer(make([]byte, 0, max))
}

func NewByte(max int) []byte {
	return make([]byte, 0, max)
}

func Benchmark_Buffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewBuffer(types.KB)
	}
}
func Benchmark_Byte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewByte(types.KB)
	}
}

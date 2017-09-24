package GoTestExample

import (
	"sync"
	"testing"
)

func Benchmark_Mutex(b *testing.B) {
	mux := new(sync.Mutex)
	for i := 0; i < b.N; i++ {
		mux.Lock()
		mux.Unlock()
	}
}

func Benchmark_Chan(b *testing.B) {
	ch := make(chan []byte, 1)
	bt := []byte("1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111")
	for i := 0; i < b.N; i++ {
		ch <- bt
		<-ch
	}
	close(ch)
}

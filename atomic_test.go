package GoTestExample

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func Test_atomic(t *testing.T) {
	var i uint64 = 0
	i = atomic.AddUint64(&i, 1)
	fmt.Println(i)
}

package GoTestExample

import (
	"bytes"
)

func NewBuffer(max int) *bytes.Buffer {
	return bytes.NewBuffer(make([]byte, 0, max))
}

func NewByte(max int) []byte {
	return make([]byte, 0, max)
}

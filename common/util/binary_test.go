package util

import (
	"testing"
)

func Test1(t *testing.T) {
	// 不好使！！
	v2 := encode("0123456789", []byte{255})
	Print(v2)
}

func Test2(t *testing.T) {
	a := 225
	a >>= 4
	a <<= 4
	Print(a)
}

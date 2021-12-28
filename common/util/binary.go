package util

import (
	"encoding/binary"
)

func Uint64ToBytes(in uint64) []byte {
	b := make([]byte, 8)
	res := make([]byte, 0)
	binary.BigEndian.PutUint64(b, in)
	for _, i := range b {
		if i == 0 {
			continue
		}
		res = append(res, i)
	}
	if len(res) == 0 {
		res = append(res, 0)
	}
	return res
}

func BytesToUint64(in []byte) uint64 {
	if len(in) < 8 {
		trim := make([]byte, 8)
		trimLen := 8 - len(in)
		for i := range trim {
			if i < trimLen {
				continue
			}
			trim[i] = in[i-trimLen]
		}
		in = trim
	}
	if len(in) > 8 {
		in = in[:8]
	}
	return binary.BigEndian.Uint64(in)
}

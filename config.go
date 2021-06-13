package goimagehash

import (
	"encoding/hex"
	"math/bits"
)

const (
	hexBits  = 4
	byteBits = 8
)

// EncodeToString ≈ hex.EncodeToString
func EncodeToString(src []byte, hexs int) string {
	opt := hex.EncodeToString(src)
	if hexs&1 == 0 {
		return opt
	}

	return opt[:len(opt)-2] + opt[len(opt)-2:len(opt)-1]
}

// HammingDistance the result will be -1 if len(a)!=len(b)
func HammingDistance(a, b []byte) int {
	if len(a) != len(b) {
		return -1
	}

	opt := 0
	for i, v := range a {
		v ^= b[i]
		opt += bits.OnesCount8(v)
	}
	return opt
}

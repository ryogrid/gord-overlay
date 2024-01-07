package model

import (
	"bytes"
	"encoding/binary"
	"math/big"
	"testing"
)

func BenchmarkPutBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		one := make([]byte, 256)
		two := make([]byte, 256)
		binary.BigEndian.PutUint64(one, 2)
		binary.BigEndian.PutUint64(two, 256)
		bytes.Compare(one, two)
	}
}

func BenchmarkBigInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		one := big.NewInt(2).Bytes()
		two := big.NewInt(256).Bytes()
		big.NewInt(0).SetBytes(one).Cmp(big.NewInt(0).SetBytes(two))
	}
}

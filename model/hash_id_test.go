package model

import (
	"math/big"
	"testing"
)

func TestNewHashID(t *testing.T) {
	addressA := "127.0.0.1:8080"
	addressB := "127.0.0.1:8081"
	a := NewHashID(addressA)
	b := NewHashID(addressA)
	c := NewHashID(addressB)
	if !a.Equals(b) {
		t.Fatalf("a is not b. a = %x, b = %x", a, b)
	}
	if a.Equals(c) {
		t.Fatalf("a must not be c. a,c = %x", a)
	}
}

func TestBetween(t *testing.T) {
	testcases := []struct {
		from     int64
		to       int64
		target   int64
		expected bool
	}{
		{
			from:     1,
			to:       3,
			target:   2,
			expected: true,
		},
		{
			from:     3,
			to:       1,
			target:   2,
			expected: false,
		},
		{
			from:     3,
			to:       1,
			target:   4,
			expected: true,
		},
	}
	for _, testcase := range testcases {
		fromID := HashID(big.NewInt(testcase.from).Bytes())
		toID := HashID(big.NewInt(testcase.to).Bytes())
		targetID := HashID(big.NewInt(testcase.target).Bytes())
		if result := targetID.Between(fromID, toID); result != testcase.expected {
			t.Fatalf("Expected %v, but actually %v. The from is %v and to is %v.", testcase.expected, result, testcase.from, testcase.to)
		}
	}
}

//func TestGreaterThanEqual(t *testing.T) {
//	tables := []struct {
//		a         HashID
//		b         HashID
//		aIsBigger bool
//	}{
//		{
//			a:         big.NewInt(2).Bytes(),
//			b:         big.NewInt(256).Bytes(),
//			aIsBigger: false,
//		},
//	}
//	for _, tc := range tables {
//		assert.Equal(t, tc.aIsBigger, tc.a.GreaterThanEqual(tc.b))
//	}
//}

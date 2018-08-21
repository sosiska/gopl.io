package popcount_test

import (
	"testing"

	"github.com/sosiska/gopl.io/ch2/popcount23"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountRange(0x1234567890ABCDEF)
	}
}

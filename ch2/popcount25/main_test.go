package popcount_test

import (
	"testing"

	"github.com/sosiska/gopl.io/ch2/popcount25"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByClearing(0x1234567890ABCDEF)
	}
}

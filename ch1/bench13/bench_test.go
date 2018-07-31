package bench13

import (
	"strings"
	"testing"
)

var t = []string{"hallo", "how", "are", "you", "lololo", "yandex"}

func BenchmarkConcat(b *testing.B) {
	var s, sep string
	for i := 0; i < b.N; i++ {
		s = ""
		for i := 0; i < len(t); i++ {
			s += sep + t[i]
			sep = " "
		}
	}
	_ = s
}

func BenchmarkRange(b *testing.B) {
	var s, sep string
	for i := 0; i < b.N; i++ {
		s = ""
		for _, v := range t {
			s += sep + v
			sep = " "
		}
	}
	_ = s
}

func BenchmarkJoin(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = ""
		s = strings.Join(t, " ")
	}
	_ = s
}

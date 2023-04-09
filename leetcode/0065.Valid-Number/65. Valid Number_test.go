package leetcode

import (
	"testing"
)

func Benchmark_Problem65(b *testing.B) {

	tcs := []struct {
		s   string
		ans bool
	}{

		{
			"0",
			true,
		},

		{
			"e",
			false,
		},

		{
			".",
			false,
		},

		{
			".1",
			true,
		},
	}
	for bbe := 0; bbe < b.N; bbe++ {
		for _, tc := range tcs {
			(isNumber(tc.s))
		}
	}
}

package leetcode

import (
	"fmt"
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

	for _, tc := range tcs {
		fmt.Printf("【input】:%v       【output】:%v\n", tc, isNumber(tc.s))
	}
}

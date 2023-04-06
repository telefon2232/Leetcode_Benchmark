package leetcode

import (
	"testing"
)

func Benchmark_Problem125(b *testing.B) {

	tcs := []struct {
		s   string
		ans bool
	}{

		{
			"0p",
			false,
		},

		{
			"0",
			true,
		},

		{
			"race a car",
			false,
		},

		{
			"A man, a plan, a canal: Panama",
			true,
		},
	}

	for _, tc := range tcs {
		(isPalindrome(tc.s))
	}
}

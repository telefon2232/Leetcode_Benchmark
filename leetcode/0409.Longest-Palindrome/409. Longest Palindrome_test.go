package leetcode

import (
	"testing"
)

type question409 struct {
	para409
	ans409
}

// para 是参数
// one 代表第一个参数
type para409 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans409 struct {
	one int
}

func Benchmark_Problem409(b *testing.B) {

	qs := []question409{

		{
			para409{"abccccdd"},
			ans409{7},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans409, q.para409
				(longestPalindrome(p.one))
			}
		}
	}
}

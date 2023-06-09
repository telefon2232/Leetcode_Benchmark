package leetcode

import (
	"testing"
)

type question32 struct {
	para32
	ans32
}

// para 是参数
// one 代表第一个参数
type para32 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans32 struct {
	one int
}

func Benchmark_Problem32(b *testing.B) {

	qs := []question32{

		{
			para32{"(()"},
			ans32{2},
		},

		{
			para32{")()())"},
			ans32{4},
		},

		{
			para32{"()(())"},
			ans32{6},
		},

		{
			para32{"()(())))"},
			ans32{6},
		},

		{
			para32{"()(()"},
			ans32{2},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans32, q.para32
				(longestValidParentheses(p.s))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question856 struct {
	para856
	ans856
}

// para 是参数
// one 代表第一个参数
type para856 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans856 struct {
	one int
}

func Benchmark_Problem856(b *testing.B) {

	qs := []question856{
		{
			para856{"()"},
			ans856{1},
		},

		{
			para856{"(())"},
			ans856{2},
		},

		{
			para856{"()()"},
			ans856{2},
		},

		{
			para856{"(()(()))"},
			ans856{6},
		},

		{
			para856{"()(())"},
			ans856{3},
		},

		{
			para856{"((()()))"},
			ans856{8},
		},
	}

	for _, q := range qs {
		_, p := q.ans856, q.para856
		(scoreOfParentheses(p.one))
	}
}

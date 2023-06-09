package leetcode

import (
	"testing"
)

type question1021 struct {
	para1021
	ans1021
}

// para 是参数
// one 代表第一个参数
type para1021 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans1021 struct {
	one string
}

func Benchmark_Problem1021(b *testing.B) {

	qs := []question1021{
		{
			para1021{"(()())(())"},
			ans1021{"()()()"},
		},

		{
			para1021{"(()())(())(()(()))"},
			ans1021{"()()()()(())"},
		},

		{
			para1021{"()()"},
			ans1021{""},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1021, q.para1021
				(removeOuterParentheses(p.one))
			}
		}
	}
}

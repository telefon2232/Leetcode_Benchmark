package leetcode

import (
	"testing"
)

type question150 struct {
	para150
	ans150
}

// para 是参数
// one 代表第一个参数
type para150 struct {
	one []string
}

// ans 是答案
// one 代表第一个答案
type ans150 struct {
	one int
}

func Benchmark_Problem150(b *testing.B) {

	qs := []question150{

		{
			para150{[]string{"18"}},
			ans150{18},
		},

		{
			para150{[]string{"2", "1", "+", "3", "*"}},
			ans150{9},
		},
		{
			para150{[]string{"4", "13", "5", "/", "+"}},
			ans150{6},
		},

		{
			para150{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}},
			ans150{22},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans150, q.para150
				(evalRPN(p.one))
			}
		}
	}
}

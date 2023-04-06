package leetcode

import (
	"testing"
)

type question682 struct {
	para682
	ans682
}

// para 是参数
// one 代表第一个参数
type para682 struct {
	one []string
}

// ans 是答案
// one 代表第一个答案
type ans682 struct {
	one int
}

func Benchmark_Problem682(b *testing.B) {

	qs := []question682{

		{
			para682{[]string{"5", "2", "C", "D", "+"}},
			ans682{30},
		},

		{
			para682{[]string{"5", "-2", "4", "C", "D", "9", "+", "+"}},
			ans682{27},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans682, q.para682
				(calPoints(p.one))
			}
		}
	}
}

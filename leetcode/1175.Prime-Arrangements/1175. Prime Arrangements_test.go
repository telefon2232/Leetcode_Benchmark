package leetcode

import (
	"testing"
)

type question1175 struct {
	para1175
	ans1175
}

// para 是参数
// one 代表第一个参数
type para1175 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans1175 struct {
	one int
}

func Benchmark_Problem1175(b *testing.B) {

	qs := []question1175{

		{
			para1175{5},
			ans1175{12},
		},

		{
			para1175{99},
			ans1175{75763854},
		},

		{
			para1175{100},
			ans1175{682289015},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1175, q.para1175
				(numPrimeArrangements(p.one))
			}
		}
	}
}

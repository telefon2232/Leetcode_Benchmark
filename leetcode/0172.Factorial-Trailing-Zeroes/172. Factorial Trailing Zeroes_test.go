package leetcode

import (
	"testing"
)

type question172 struct {
	para172
	ans172
}

// para 是参数
// one 代表第一个参数
type para172 struct {
	s int
}

// ans 是答案
// one 代表第一个答案
type ans172 struct {
	one int
}

func Benchmark_Problem172(b *testing.B) {

	qs := []question172{

		{
			para172{3},
			ans172{0},
		},

		{
			para172{5},
			ans172{1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans172, q.para172
				(trailingZeroes(p.s))
			}
		}
	}
}

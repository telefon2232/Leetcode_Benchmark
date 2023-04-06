package leetcode

import (
	"testing"
)

type question343 struct {
	para343
	ans343
}

// para 是参数
// one 代表第一个参数
type para343 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans343 struct {
	one int
}

func Benchmark_Problem343(b *testing.B) {

	qs := []question343{

		{
			para343{2},
			ans343{1},
		},

		{
			para343{10},
			ans343{36},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans343, q.para343
				(integerBreak(p.one))
			}
		}
	}
}

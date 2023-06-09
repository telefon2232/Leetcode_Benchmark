package leetcode

import (
	"testing"
)

type question279 struct {
	para279
	ans279
}

// para 是参数
// one 代表第一个参数
type para279 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans279 struct {
	one int
}

func Benchmark_Problem279(b *testing.B) {

	qs := []question279{

		{
			para279{13},
			ans279{2},
		},
		{
			para279{12},
			ans279{3},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans279, q.para279
				(numSquares(p.n))
			}
		}
	}
}

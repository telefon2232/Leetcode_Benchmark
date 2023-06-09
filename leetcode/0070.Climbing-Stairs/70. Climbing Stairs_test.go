package leetcode

import (
	"testing"
)

type question70 struct {
	para70
	ans70
}

// para 是参数
// one 代表第一个参数
type para70 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans70 struct {
	one int
}

func Benchmark_Problem70(b *testing.B) {

	qs := []question70{

		{
			para70{2},
			ans70{2},
		},

		{
			para70{3},
			ans70{3},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans70, q.para70
				(climbStairs(p.n))
			}
		}
	}
}

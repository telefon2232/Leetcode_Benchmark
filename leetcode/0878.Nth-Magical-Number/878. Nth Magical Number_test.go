package leetcode

import (
	"testing"
)

type question878 struct {
	para878
	ans878
}

// para 是参数
// one 代表第一个参数
type para878 struct {
	n int
	a int
	b int
}

// ans 是答案
// one 代表第一个答案
type ans878 struct {
	one int
}

func Benchmark_Problem878(b *testing.B) {

	qs := []question878{

		{
			para878{1, 2, 3},
			ans878{2},
		},

		{
			para878{4, 2, 3},
			ans878{6},
		},

		{
			para878{5, 2, 4},
			ans878{10},
		},

		{
			para878{3, 6, 4},
			ans878{8},
		},

		{
			para878{1000000000, 40000, 40000},
			ans878{999720007},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans878, q.para878
				(nthMagicalNumber(p.n, p.a, p.b))
			}
		}
	}
}

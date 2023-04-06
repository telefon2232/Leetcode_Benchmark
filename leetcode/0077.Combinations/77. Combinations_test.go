package leetcode

import (
	"testing"
)

type question77 struct {
	para77
	ans77
}

// para 是参数
// one 代表第一个参数
type para77 struct {
	n int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans77 struct {
	one [][]int
}

func Benchmark_Problem77(b *testing.B) {

	qs := []question77{

		{
			para77{4, 2},
			ans77{[][]int{{2, 4}, {3, 4}, {2, 3}, {1, 2}, {1, 3}, {1, 4}}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans77, q.para77
				(combine(p.n, p.k))
			}
		}
	}
}

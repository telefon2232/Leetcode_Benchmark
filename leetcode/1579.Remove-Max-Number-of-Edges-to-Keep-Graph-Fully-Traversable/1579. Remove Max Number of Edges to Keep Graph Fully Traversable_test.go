package leetcode

import (
	"testing"
)

type question1579 struct {
	para1579
	ans1579
}

// para 是参数
// one 代表第一个参数
type para1579 struct {
	n     int
	edges [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1579 struct {
	one int
}

func Benchmark_Problem1579(b *testing.B) {

	qs := []question1579{

		{
			para1579{4, [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 3}, {1, 2, 4}, {1, 1, 2}, {2, 3, 4}}},
			ans1579{2},
		},

		{
			para1579{4, [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 4}, {2, 1, 4}}},
			ans1579{0},
		},

		{
			para1579{4, [][]int{{3, 2, 3}, {1, 1, 2}, {2, 3, 4}}},
			ans1579{-1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1579, q.para1579
				(maxNumEdgesToRemove(p.n, p.edges))
			}
		}
	}
}

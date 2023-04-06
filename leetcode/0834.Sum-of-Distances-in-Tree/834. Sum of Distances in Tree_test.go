package leetcode

import (
	"testing"
)

type question834 struct {
	para834
	ans834
}

// para 是参数
// one 代表第一个参数
type para834 struct {
	N     int
	edges [][]int
}

// ans 是答案
// one 代表第一个答案
type ans834 struct {
	one []int
}

func Benchmark_Problem834(b *testing.B) {

	qs := []question834{

		{
			para834{4, [][]int{{1, 2}, {3, 2}, {3, 0}}},
			ans834{[]int{6, 6, 4, 4}},
		},

		{
			para834{6, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}},
			ans834{[]int{8, 12, 6, 10, 10, 10}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans834, q.para834
				(sumOfDistancesInTree(p.N, p.edges))
			}
		}
	}
}

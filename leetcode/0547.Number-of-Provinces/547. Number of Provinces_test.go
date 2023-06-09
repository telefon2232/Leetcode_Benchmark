package leetcode

import (
	"testing"
)

type question547 struct {
	para547
	ans547
}

// para 是参数
// one 代表第一个参数
type para547 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans547 struct {
	one int
}

func Benchmark_Problem547(b *testing.B) {

	qs := []question547{

		{
			para547{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
			ans547{3},
		},

		{
			para547{[][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}},
			ans547{2},
		},

		{
			para547{[][]int{{1, 1, 0}, {1, 1, 1}, {0, 1, 1}}},
			ans547{1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans547, q.para547
				(findCircleNum(p.one))
			}
		}
	}
}

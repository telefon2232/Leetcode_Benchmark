package leetcode

import (
	"testing"
)

type question807 struct {
	para807
	ans807
}

// para 是参数
type para807 struct {
	grid [][]int
}

// ans 是答案
type ans807 struct {
	ans int
}

func Benchmark_Problem807(b *testing.B) {

	qs := []question807{

		{
			para807{[][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}},
			ans807{35},
		},

		{
			para807{[][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
			ans807{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans807, q.para807
				(maxIncreaseKeepingSkyline(p.grid))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question1030 struct {
	para1030
	ans1030
}

// para 是参数
// one 代表第一个参数
type para1030 struct {
	R  int
	C  int
	r0 int
	c0 int
}

// ans 是答案
// one 代表第一个答案
type ans1030 struct {
	one [][]int
}

func Benchmark_Problem1030(b *testing.B) {

	qs := []question1030{

		{
			para1030{1, 2, 0, 0},
			ans1030{[][]int{{0, 0}, {0, 1}}},
		},

		{
			para1030{2, 2, 0, 1},
			ans1030{[][]int{{0, 1}, {0, 0}, {1, 1}, {1, 0}}},
		},

		{
			para1030{2, 3, 1, 2},
			ans1030{[][]int{{1, 2}, {0, 2}, {1, 1}, {0, 1}, {1, 0}, {0, 0}}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1030, q.para1030
				(allCellsDistOrder(p.R, p.C, p.r0, p.c0))
			}
		}
	}
}

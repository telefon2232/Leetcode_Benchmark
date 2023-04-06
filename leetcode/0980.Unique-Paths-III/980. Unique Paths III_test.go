package leetcode

import (
	"testing"
)

type question980 struct {
	para980
	ans980
}

// para 是参数
// one 代表第一个参数
type para980 struct {
	grid [][]int
}

// ans 是答案
// one 代表第一个答案
type ans980 struct {
	one int
}

func Benchmark_Problem980(b *testing.B) {

	qs := []question980{

		{
			para980{[][]int{
				{1, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 2, -1},
			}},
			ans980{2},
		},

		{
			para980{[][]int{
				{1, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 2},
			}},
			ans980{4},
		},

		{
			para980{[][]int{
				{1, 0},
				{0, 2},
			}},
			ans980{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans980, q.para980
				(uniquePathsIII(p.grid))
			}
		}
	}
}

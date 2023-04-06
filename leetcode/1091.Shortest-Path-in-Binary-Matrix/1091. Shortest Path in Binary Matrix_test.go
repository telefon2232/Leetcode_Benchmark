package leetcode

import (
	"testing"
)

type question1091 struct {
	para1091
	ans1091
}

// para 是参数
// one 代表第一个参数
type para1091 struct {
	grid [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1091 struct {
	one int
}

func Benchmark_Problem1091(b *testing.B) {

	qs := []question1091{

		{
			para1091{[][]int{{0, 1}, {1, 0}}},
			ans1091{2},
		},

		{
			para1091{[][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}},
			ans1091{4},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1091, q.para1091
				(shortestPathBinaryMatrix(p.grid))

			}
		}
	}
}

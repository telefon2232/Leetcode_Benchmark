package leetcode

import (
	"testing"
)

type question812 struct {
	para812
	ans812
}

// para 是参数
// one 代表第一个参数
type para812 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans812 struct {
	one float64
}

func Benchmark_Problem812(b *testing.B) {

	qs := []question812{

		{
			para812{[][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}},
			ans812{2.0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans812, q.para812
				(largestTriangleArea(p.one))
			}
		}
	}
}

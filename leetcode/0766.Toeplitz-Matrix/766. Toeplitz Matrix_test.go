package leetcode

import (
	"testing"
)

type question766 struct {
	para766
	ans766
}

// para 是参数
// one 代表第一个参数
type para766 struct {
	A [][]int
}

// ans 是答案
// one 代表第一个答案
type ans766 struct {
	B bool
}

func Benchmark_Problem766(b *testing.B) {

	qs := []question766{

		{
			para766{[][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}},
			ans766{true},
		},

		{
			para766{[][]int{{1, 2}, {2, 2}}},
			ans766{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans766, q.para766
				(isToeplitzMatrix(p.A))
			}
		}
	}
}

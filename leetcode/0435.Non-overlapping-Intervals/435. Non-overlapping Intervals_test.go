package leetcode

import (
	"testing"
)

type question435 struct {
	para435
	ans435
}

// para 是参数
// one 代表第一个参数
type para435 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans435 struct {
	one int
}

func Benchmark_Problem435(b *testing.B) {

	qs := []question435{

		{
			para435{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}},
			ans435{1},
		},

		{
			para435{[][]int{{1, 2}, {1, 2}, {1, 2}}},
			ans435{2},
		},

		{
			para435{[][]int{{1, 2}, {2, 3}}},
			ans435{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans435, q.para435
				(eraseOverlapIntervals1(p.one))
			}
		}
	}
}

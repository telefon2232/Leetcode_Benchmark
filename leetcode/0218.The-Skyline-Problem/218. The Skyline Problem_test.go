package leetcode

import (
	"testing"
)

type question218 struct {
	para218
	ans218
}

// para 是参数
// one 代表第一个参数
type para218 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans218 struct {
	one [][]int
}

func Benchmark_Problem218(b *testing.B) {

	qs := []question218{

		{
			para218{[][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}},
			ans218{[][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}}},
		},

		{
			para218{[][]int{{1, 2, 1}, {1, 2, 2}, {1, 2, 3}, {2, 3, 1}, {2, 3, 2}, {2, 3, 3}}},
			ans218{[][]int{{1, 3}, {3, 0}}},
		},

		{
			para218{[][]int{{4, 9, 10}, {4, 9, 15}, {4, 9, 12}, {10, 12, 10}, {10, 12, 8}}},
			ans218{[][]int{{4, 15}, {9, 0}, {10, 10}, {12, 0}}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans218, q.para218
		(getSkyline(p.one))
	}
}}}

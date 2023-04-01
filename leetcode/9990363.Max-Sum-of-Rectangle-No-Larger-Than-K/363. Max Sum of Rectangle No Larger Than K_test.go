package leetcode

import (
	"testing"
)

type question363 struct {
	para363
	ans363
}

// para 是参数
// one 代表第一个参数
type para363 struct {
	one [][]int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans363 struct {
	one int
}

func Benchmark_Problem363(b *testing.B) {

	qs := []question363{

		{
			para363{[][]int{{1, 0, 1}, {0, -2, 3}}, 2},
			ans363{2},
		},

		{
			para363{[][]int{{2, 2, -1}}, 0},
			ans363{-1},
		},

		{
			para363{[][]int{{5, -4, -3, 4}, {-3, -4, 4, 5}, {5, 1, 5, -4}}, 8},
			ans363{8},
		},

		{
			para363{[][]int{{5, -4, -3, 4}, {-3, -4, 4, 5}, {5, 1, 5, -4}}, 3},
			ans363{2},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans363, q.para363
		(maxSumSubmatrix(p.one, p.k))
	}
}}}

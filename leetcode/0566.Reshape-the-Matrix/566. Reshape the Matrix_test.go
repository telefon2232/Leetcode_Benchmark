package leetcode

import (
	"testing"
)

type question566 struct {
	para566
	ans566
}

// para 是参数
// one 代表第一个参数
type para566 struct {
	nums [][]int
	r    int
	c    int
}

// ans 是答案
// one 代表第一个答案
type ans566 struct {
	one [][]int
}

func Benchmark_Problem566(b *testing.B) {

	qs := []question566{

		{
			para566{[][]int{{1, 2}, {3, 4}}, 1, 4},
			ans566{[][]int{{1, 2, 3, 4}}},
		},

		{
			para566{[][]int{{1, 2}, {3, 4}}, 2, 4},
			ans566{[][]int{{1, 2}, {3, 4}}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans566, q.para566
		(matrixReshape(p.nums, p.r, p.c))
	}
}}}

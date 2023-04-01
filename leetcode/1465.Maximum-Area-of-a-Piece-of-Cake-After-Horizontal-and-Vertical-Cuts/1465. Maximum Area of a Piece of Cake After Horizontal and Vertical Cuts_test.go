package leetcode

import (
	"testing"
)

type question1465 struct {
	para1465
	ans1465
}

// para 是参数
// one 代表第一个参数
type para1465 struct {
	h              int
	w              int
	horizontalCuts []int
	verticalCuts   []int
}

// ans 是答案
// one 代表第一个答案
type ans1465 struct {
	one int
}

func Benchmark_Problem1465(b *testing.B) {

	qs := []question1465{

		{
			para1465{5, 4, []int{1, 2, 4}, []int{1, 3}},
			ans1465{4},
		},

		{
			para1465{5, 4, []int{3, 1}, []int{1}},
			ans1465{6},
		},

		{
			para1465{5, 4, []int{3}, []int{3}},
			ans1465{9},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1465, q.para1465
				(maxArea(p.h, p.w, p.horizontalCuts, p.verticalCuts))
			}
		}
	}
}

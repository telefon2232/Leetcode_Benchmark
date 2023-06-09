package leetcode

import (
	"testing"
)

type question11 struct {
	para11
	ans11
}

// para 是参数
// one 代表第一个参数
type para11 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans11 struct {
	one int
}

func Benchmark_Problem11(b *testing.B) {

	qs := []question11{

		{
			para11{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			ans11{49},
		},

		{
			para11{[]int{1, 1}},
			ans11{1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans11, q.para11
				(maxArea(p.one))
			}
		}
	}
}

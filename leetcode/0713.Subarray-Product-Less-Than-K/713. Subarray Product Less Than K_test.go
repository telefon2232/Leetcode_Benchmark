package leetcode

import (
	"testing"
)

type question713 struct {
	para713
	ans713
}

// para 是参数
// one 代表第一个参数
type para713 struct {
	s []int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans713 struct {
	one int
}

func Benchmark_Problem713(b *testing.B) {

	qs := []question713{

		{
			para713{[]int{10, 5, 2, 6}, 100},
			ans713{8},
		},

		{
			para713{[]int{10, 9, 10, 4, 3, 8, 3, 3, 6, 2, 10, 10, 9, 3}, 19},
			ans713{18},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans713, q.para713
		(numSubarrayProductLessThanK(p.s, p.k))
	}
}}}

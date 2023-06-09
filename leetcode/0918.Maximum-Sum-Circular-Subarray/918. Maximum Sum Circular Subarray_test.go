package leetcode

import (
	"testing"
)

type question918 struct {
	para918
	ans918
}

// para 是参数
// one 代表第一个参数
type para918 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans918 struct {
	one int
}

func Benchmark_Problem918(b *testing.B) {

	qs := []question918{

		{
			para918{[]int{1, -2, 3, -2}},
			ans918{3},
		},

		{
			para918{[]int{5, -3, 5}},
			ans918{10},
		},

		{
			para918{[]int{3, -1, 2, -1}},
			ans918{4},
		},

		{
			para918{[]int{3, -2, 2, -3}},
			ans918{3},
		},

		{
			para918{[]int{-2, -3, -1}},
			ans918{-1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans918, q.para918
				(maxSubarraySumCircular(p.one))
			}
		}
	}
}

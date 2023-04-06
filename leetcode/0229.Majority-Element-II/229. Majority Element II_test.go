package leetcode

import (
	"testing"
)

type question229 struct {
	para229
	ans229
}

// para 是参数
// one 代表第一个参数
type para229 struct {
	s []int
}

// ans 是答案
// one 代表第一个答案
type ans229 struct {
	one []int
}

func Benchmark_Problem229(b *testing.B) {

	qs := []question229{

		{
			para229{[]int{3, 2, 3}},
			ans229{[]int{3}},
		},

		{
			para229{[]int{1, 1, 1, 3, 3, 2, 2, 2}},
			ans229{[]int{1, 2}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans229, q.para229
				(majorityElement229(p.s))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question746 struct {
	para746
	ans746
}

// para 是参数
// one 代表第一个参数
type para746 struct {
	c []int
}

// ans 是答案
// one 代表第一个答案
type ans746 struct {
	one int
}

func Benchmark_Problem746(b *testing.B) {

	qs := []question746{

		{
			para746{[]int{10, 15, 20}},
			ans746{15},
		},

		{
			para746{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}},
			ans746{6},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans746, q.para746
				(minCostClimbingStairs(p.c))
			}
		}
	}
}

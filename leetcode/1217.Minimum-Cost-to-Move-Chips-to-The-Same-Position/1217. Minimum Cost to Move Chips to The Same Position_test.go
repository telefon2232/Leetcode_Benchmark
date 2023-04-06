package leetcode

import (
	"testing"
)

type question1217 struct {
	para1217
	ans1217
}

// para 是参数
// one 代表第一个参数
type para1217 struct {
	arr []int
}

// ans 是答案
// one 代表第一个答案
type ans1217 struct {
	one int
}

func Benchmark_Problem1217(b *testing.B) {

	qs := []question1217{

		{
			para1217{[]int{1, 2, 3}},
			ans1217{1},
		},

		{
			para1217{[]int{2, 2, 2, 3, 3}},
			ans1217{2},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1217, q.para1217
				(minCostToMoveChips(p.arr))
			}
		}
	}
}

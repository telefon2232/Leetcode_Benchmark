package leetcode

import (
	"testing"
)

type question525 struct {
	para525
	ans525
}

// para 是参数
// one 代表第一个参数
type para525 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans525 struct {
	one int
}

func Benchmark_Problem525(b *testing.B) {

	qs := []question525{

		{
			para525{[]int{0, 1}},
			ans525{2},
		},

		{
			para525{[]int{0, 1, 0}},
			ans525{2},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans525, q.para525
				(findMaxLength(p.nums))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question309 struct {
	para309
	ans309
}

// para 是参数
// one 代表第一个参数
type para309 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans309 struct {
	one int
}

func Benchmark_Problem309(b *testing.B) {

	qs := []question309{

		{
			para309{[]int{}},
			ans309{0},
		},

		{
			para309{[]int{2, 1, 4, 5, 2, 9, 7}},
			ans309{10},
		},

		{
			para309{[]int{6, 1, 3, 2, 4, 7}},
			ans309{6},
		},

		{
			para309{[]int{1, 2, 3, 0, 2}},
			ans309{3},
		},

		{
			para309{[]int{7, 1, 5, 3, 6, 4}},
			ans309{5},
		},

		{
			para309{[]int{7, 6, 4, 3, 1}},
			ans309{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans309, q.para309
				(maxProfit309(p.one))
			}
		}
	}
}

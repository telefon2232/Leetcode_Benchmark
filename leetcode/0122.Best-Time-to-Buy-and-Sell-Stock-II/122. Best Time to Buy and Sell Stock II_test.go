package leetcode

import (
	"testing"
)

type question122 struct {
	para122
	ans122
}

// para 是参数
// one 代表第一个参数
type para122 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans122 struct {
	one int
}

func Benchmark_Problem122(b *testing.B) {

	qs := []question122{

		{
			para122{[]int{}},
			ans122{0},
		},

		{
			para122{[]int{7, 1, 5, 3, 6, 4}},
			ans122{7},
		},

		{
			para122{[]int{7, 6, 4, 3, 1}},
			ans122{0},
		},

		{
			para122{[]int{1, 2, 3, 4, 5}},
			ans122{4},
		},

		{
			para122{[]int{1, 2, 10, 11, 12, 15}},
			ans122{14},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans122, q.para122
				(maxProfit122(p.one))
			}
		}
	}
}

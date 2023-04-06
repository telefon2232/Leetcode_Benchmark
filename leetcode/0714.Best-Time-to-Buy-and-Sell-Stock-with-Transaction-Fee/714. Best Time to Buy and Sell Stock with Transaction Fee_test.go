package leetcode

import (
	"testing"
)

type question714 struct {
	para714
	ans714
}

// para 是参数
// one 代表第一个参数
type para714 struct {
	one []int
	f   int
}

// ans 是答案
// one 代表第一个答案
type ans714 struct {
	one int
}

func Benchmark_Problem714(b *testing.B) {

	qs := []question714{

		{
			para714{[]int{}, 0},
			ans714{0},
		},

		{
			para714{[]int{7, 1, 5, 3, 6, 4}, 0},
			ans714{7},
		},

		{
			para714{[]int{7, 6, 4, 3, 1}, 0},
			ans714{0},
		},

		{
			para714{[]int{1, 3, 2, 8, 4, 9}, 2},
			ans714{8},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans714, q.para714
				(maxProfit714(p.one, p.f))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question560 struct {
	para560
	ans560
}

// para 是参数
// one 代表第一个参数
type para560 struct {
	nums []int
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans560 struct {
	one int
}

func Benchmark_Problem560(b *testing.B) {

	qs := []question560{

		{
			para560{[]int{1, 1, 1}, 2},
			ans560{2},
		},

		{
			para560{[]int{1, 2, 3}, 3},
			ans560{2},
		},

		{
			para560{[]int{1}, 0},
			ans560{0},
		},

		{
			para560{[]int{-1, -1, 1}, 0},
			ans560{1},
		},

		{
			para560{[]int{1, -1, 0}, 0},
			ans560{3},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans560, q.para560
				(subarraySum(p.nums, p.k))
			}
		}
	}
}

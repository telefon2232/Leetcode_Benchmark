package leetcode

import (
	"testing"
)

type question930 struct {
	para930
	ans930
}

// para 是参数
// one 代表第一个参数
type para930 struct {
	s []int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans930 struct {
	one int
}

func Benchmark_Problem930(b *testing.B) {

	qs := []question930{

		{
			para930{[]int{1, 0, 1, 0, 1}, 2},
			ans930{4},
		},

		{
			para930{[]int{0, 0, 0, 0, 0}, 0},
			ans930{15},
		},

		{
			para930{[]int{1, 0, 1, 1, 1, 1, 0, 1, 0, 1}, 2},
			ans930{4},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans930, q.para930
				(numSubarraysWithSum(p.s, p.k))
			}
		}
	}
}

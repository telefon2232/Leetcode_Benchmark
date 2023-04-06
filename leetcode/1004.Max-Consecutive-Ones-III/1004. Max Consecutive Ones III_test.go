package leetcode

import (
	"testing"
)

type question1004 struct {
	para1004
	ans1004
}

// para 是参数
// one 代表第一个参数
type para1004 struct {
	s []int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans1004 struct {
	one int
}

func Benchmark_Problem1004(b *testing.B) {

	qs := []question1004{

		{
			para1004{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2},
			ans1004{6},
		},

		{
			para1004{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3},
			ans1004{10},
		},

		{
			para1004{[]int{0, 0, 0, 1}, 4},
			ans1004{4},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1004, q.para1004
				(longestOnes(p.s, p.k))
			}
		}
	}
}

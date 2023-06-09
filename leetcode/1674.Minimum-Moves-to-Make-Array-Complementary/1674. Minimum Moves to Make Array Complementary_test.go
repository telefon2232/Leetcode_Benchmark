package leetcode

import (
	"testing"
)

type question1674 struct {
	para1674
	ans1674
}

// para 是参数
// one 代表第一个参数
type para1674 struct {
	nums  []int
	limit int
}

// ans 是答案
// one 代表第一个答案
type ans1674 struct {
	one int
}

func Benchmark_Problem1674(b *testing.B) {

	qs := []question1674{

		{
			para1674{[]int{1, 2, 4, 3}, 4},
			ans1674{1},
		},

		{
			para1674{[]int{1, 2, 2, 1}, 2},
			ans1674{2},
		},

		{
			para1674{[]int{1, 2, 1, 2}, 2},
			ans1674{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1674, q.para1674
				(minMoves(p.nums, p.limit))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question1437 struct {
	para1437
	ans1437
}

// para 是参数
// one 代表第一个参数
type para1437 struct {
	nums []int
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans1437 struct {
	one bool
}

func Benchmark_Problem1437(b *testing.B) {

	qs := []question1437{

		{
			para1437{[]int{1, 0, 0, 0, 1, 0, 0, 1}, 2},
			ans1437{true},
		},

		{
			para1437{[]int{1, 0, 0, 1, 0, 1}, 2},
			ans1437{false},
		},

		{
			para1437{[]int{1, 1, 1, 1, 1}, 0},
			ans1437{true},
		},

		{
			para1437{[]int{0, 1, 0, 1}, 1},
			ans1437{true},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1437, q.para1437

				(kLengthApart(p.nums, p.k))
			}
		}
	}
}

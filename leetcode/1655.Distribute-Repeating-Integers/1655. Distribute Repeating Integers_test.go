package leetcode

import (
	"testing"
)

type question1655 struct {
	para1655
	ans1655
}

// para 是参数
// one 代表第一个参数
type para1655 struct {
	nums     []int
	quantity []int
}

// ans 是答案
// one 代表第一个答案
type ans1655 struct {
	one bool
}

func Benchmark_Problem1655(b *testing.B) {

	qs := []question1655{

		{
			para1655{[]int{1, 2, 3, 4}, []int{2}},
			ans1655{false},
		},

		{
			para1655{[]int{1, 2, 3, 3}, []int{2}},
			ans1655{true},
		},

		{
			para1655{[]int{1, 1, 2, 2}, []int{2, 2}},
			ans1655{true},
		},

		{
			para1655{[]int{1, 1, 2, 3}, []int{2, 2}},
			ans1655{false},
		},

		{
			para1655{[]int{1, 1, 1, 1, 1}, []int{2, 3}},
			ans1655{true},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1655, q.para1655
				(canDistribute(p.nums, p.quantity))
			}
		}
	}
}

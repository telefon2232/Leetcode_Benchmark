package leetcode

import (
	"testing"
)

type question1480 struct {
	para1480
	ans1480
}

// para 是参数
// one 代表第一个参数
type para1480 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans1480 struct {
	one []int
}

func Benchmark_Problem1480(b *testing.B) {

	qs := []question1480{

		{
			para1480{[]int{1, 2, 3, 4}},
			ans1480{[]int{1, 3, 6, 10}},
		},

		{
			para1480{[]int{1, 1, 1, 1, 1}},
			ans1480{[]int{1, 2, 3, 4, 5}},
		},

		{
			para1480{[]int{3, 1, 2, 10, 1}},
			ans1480{[]int{3, 4, 6, 16, 17}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1480, q.para1480
				(runningSum(p.nums))
			}
		}
	}
}

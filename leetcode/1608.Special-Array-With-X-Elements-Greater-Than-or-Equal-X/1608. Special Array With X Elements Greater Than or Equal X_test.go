package leetcode

import (
	"testing"
)

type question1608 struct {
	para1608
	ans1608
}

// para 是参数
// one 代表第一个参数
type para1608 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans1608 struct {
	one int
}

func Benchmark_Problem1608(b *testing.B) {

	qs := []question1608{

		{
			para1608{[]int{3, 5}},
			ans1608{2},
		},

		{
			para1608{[]int{0, 0}},
			ans1608{-1},
		},

		{
			para1608{[]int{0, 4, 3, 0, 4}},
			ans1608{3},
		},

		{
			para1608{[]int{3, 6, 7, 7, 0}},
			ans1608{-1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1608, q.para1608
				(specialArray(p.nums))
			}
		}
	}
}

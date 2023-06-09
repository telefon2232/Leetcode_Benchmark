package leetcode

import (
	"testing"
)

type question1695 struct {
	para1695
	ans1695
}

// para 是参数
// one 代表第一个参数
type para1695 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans1695 struct {
	one int
}

func Benchmark_Problem1695(b *testing.B) {

	qs := []question1695{

		{
			para1695{[]int{4, 2, 4, 5, 6}},
			ans1695{17},
		},

		{
			para1695{[]int{5, 2, 1, 2, 5, 2, 1, 2, 5}},
			ans1695{8},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1695, q.para1695
				(maximumUniqueSubarray(p.nums))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question1748 struct {
	para1748
	ans1748
}

// para 是参数
// one 代表第一个参数
type para1748 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans1748 struct {
	one int
}

func Benchmark_Problem1748(b *testing.B) {

	qs := []question1748{

		{
			para1748{[]int{1, 2, 3, 2}},
			ans1748{4},
		},

		{
			para1748{[]int{1, 1, 1, 1, 1}},
			ans1748{0},
		},

		{
			para1748{[]int{1, 2, 3, 4, 5}},
			ans1748{15},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1748, q.para1748
				(sumOfUnique(p.nums))
			}
		}
	}
}

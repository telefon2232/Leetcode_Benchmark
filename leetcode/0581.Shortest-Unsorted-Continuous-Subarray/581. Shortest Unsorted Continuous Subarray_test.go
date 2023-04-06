package leetcode

import (
	"testing"
)

type question581 struct {
	para581
	ans581
}

// para 是参数
// one 代表第一个参数
type para581 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans581 struct {
	one int
}

func Benchmark_Problem581(b *testing.B) {

	qs := []question581{

		{
			para581{[]int{2, 6, 4, 8, 10, 9, 15}},
			ans581{5},
		},

		{
			para581{[]int{1, 2, 3, 4}},
			ans581{0},
		},

		{
			para581{[]int{1}},
			ans581{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans581, q.para581
				(findUnsortedSubarray(p.nums))
			}
		}
	}
}

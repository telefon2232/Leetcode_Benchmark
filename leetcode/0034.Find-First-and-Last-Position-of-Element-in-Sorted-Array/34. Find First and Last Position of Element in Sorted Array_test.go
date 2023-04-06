package leetcode

import (
	"testing"
)

type question34 struct {
	para34
	ans34
}

// para 是参数
// one 代表第一个参数
type para34 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans34 struct {
	one []int
}

func Benchmark_Problem34(b *testing.B) {

	qs := []question34{

		{
			para34{[]int{5, 7, 7, 8, 8, 10}, 8},
			ans34{[]int{3, 4}},
		},

		{
			para34{[]int{5, 7, 7, 8, 8, 10}, 6},
			ans34{[]int{-1, -1}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans34, q.para34
				(searchRange(p.nums, p.target))
			}
		}
	}
}

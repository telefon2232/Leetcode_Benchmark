package leetcode

import (
	"testing"
)

type question724 struct {
	para724
	ans724
}

// para 是参数
// one 代表第一个参数
type para724 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans724 struct {
	n int
}

func Benchmark_Problem724(b *testing.B) {

	qs := []question724{

		{
			para724{[]int{1, 7, 3, 6, 5, 6}},
			ans724{3},
		},

		{
			para724{[]int{1, 2, 3}},
			ans724{-1},
		},

		{
			para724{[]int{-1, -1, -1, -1, -1, -1}},
			ans724{-1},
		},

		{
			para724{[]int{-1, -1, -1, -1, -1, 0}},
			ans724{2},
		},

		{
			para724{[]int{1}},
			ans724{0},
		},

		{
			para724{[]int{}},
			ans724{-1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans724, q.para724
				(pivotIndex(p.nums))
			}
		}
	}
}

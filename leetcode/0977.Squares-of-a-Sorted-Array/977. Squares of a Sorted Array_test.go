package leetcode

import (
	"testing"
)

type question977 struct {
	para977
	ans977
}

// para 是参数
// one 代表第一个参数
type para977 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans977 struct {
	one []int
}

func Benchmark_Problem977(b *testing.B) {

	qs := []question977{

		{
			para977{[]int{-4, -1, 0, 3, 10}},
			ans977{[]int{0, 1, 9, 16, 100}},
		},

		{
			para977{[]int{1}},
			ans977{[]int{1}},
		},

		{
			para977{[]int{-7, -3, 2, 3, 11}},
			ans977{[]int{4, 9, 9, 49, 121}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans977, q.para977
				(sortedSquares(p.one))
			}
		}
	}
}

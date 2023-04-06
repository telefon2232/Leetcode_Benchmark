package leetcode

import (
	"testing"
)

type question989 struct {
	para989
	ans989
}

// para 是参数
// one 代表第一个参数
type para989 struct {
	A []int
	K int
}

// ans 是答案
// one 代表第一个答案
type ans989 struct {
	one []int
}

func Benchmark_Problem989(b *testing.B) {

	qs := []question989{

		{
			para989{[]int{1, 2, 0, 0}, 34},
			ans989{[]int{1, 2, 3, 4}},
		},

		{
			para989{[]int{2, 7, 4}, 181},
			ans989{[]int{4, 5, 5}},
		},

		{
			para989{[]int{2, 1, 5}, 806},
			ans989{[]int{1, 0, 2, 1}},
		},

		{
			para989{[]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 1},
			ans989{[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans989, q.para989
				(addToArrayForm(p.A, p.K))
			}
		}
	}
}

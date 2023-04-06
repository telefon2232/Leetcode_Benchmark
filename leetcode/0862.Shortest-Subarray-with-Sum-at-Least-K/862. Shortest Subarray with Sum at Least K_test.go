package leetcode

import (
	"testing"
)

type question862 struct {
	para862
	ans862
}

// para 是参数
// one 代表第一个参数
type para862 struct {
	A []int
	K int
}

// ans 是答案
// one 代表第一个答案
type ans862 struct {
	one int
}

func Benchmark_Problem862(b *testing.B) {

	qs := []question862{
		{
			para862{[]int{1}, 1},
			ans862{1},
		},

		{
			para862{[]int{1, 2}, 4},
			ans862{-1},
		},

		{
			para862{[]int{2, -1, 2}, 3},
			ans862{3},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans862, q.para862
				(shortestSubarray(p.A, p.K))
			}
		}
	}
}

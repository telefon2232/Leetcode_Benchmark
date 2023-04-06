package leetcode

import (
	"testing"
)

type question1025 struct {
	para1025
	ans1025
}

// para 是参数
// one 代表第一个参数
type para1025 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans1025 struct {
	one bool
}

func Benchmark_Problem1025(b *testing.B) {

	qs := []question1025{
		{
			para1025{2},
			ans1025{true},
		},

		{
			para1025{3},
			ans1025{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1025, q.para1025
				(divisorGame(p.one))
			}
		}
	}
}

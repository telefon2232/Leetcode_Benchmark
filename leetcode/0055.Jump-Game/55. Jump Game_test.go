package leetcode

import (
	"testing"
)

type question55 struct {
	para55
	ans55
}

// para 是参数
// one 代表第一个参数
type para55 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans55 struct {
	one bool
}

func Benchmark_Problem55(b *testing.B) {

	qs := []question55{

		{
			para55{[]int{2, 3, 1, 1, 4}},
			ans55{true},
		},
		{
			para55{[]int{3, 2, 1, 0, 4}},
			ans55{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans55, q.para55
				(canJump(p.one))
			}
		}
	}
}

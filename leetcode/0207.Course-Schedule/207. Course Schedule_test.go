package leetcode

import (
	"testing"
)

type question207 struct {
	para207
	ans207
}

// para 是参数
// one 代表第一个参数
type para207 struct {
	one int
	pre [][]int
}

// ans 是答案
// one 代表第一个答案
type ans207 struct {
	one bool
}

func Benchmark_Problem207(b *testing.B) {

	qs := []question207{

		{
			para207{2, [][]int{{1, 0}}},
			ans207{true},
		},

		{
			para207{2, [][]int{{1, 0}, {0, 1}}},
			ans207{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans207, q.para207
				(canFinish(p.one, p.pre))
			}
		}
	}
}

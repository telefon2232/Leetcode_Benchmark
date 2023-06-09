package leetcode

import (
	"testing"
)

type question338 struct {
	para338
	ans338
}

// para 是参数
// one 代表第一个参数
type para338 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans338 struct {
	one []int
}

func Benchmark_Problem338(b *testing.B) {

	qs := []question338{

		{
			para338{2},
			ans338{[]int{0, 1, 1}},
		},

		{
			para338{5},
			ans338{[]int{0, 1, 1, 2, 1, 2}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans338, q.para338
				(countBits(p.one))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question946 struct {
	para946
	ans946
}

// para 是参数
// one 代表第一个参数
type para946 struct {
	one []int
	two []int
}

// ans 是答案
// one 代表第一个答案
type ans946 struct {
	one bool
}

func Benchmark_Problem946(b *testing.B) {

	qs := []question946{
		{
			para946{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
			ans946{true},
		},

		{
			para946{[]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}},
			ans946{false},
		},

		{
			para946{[]int{1, 0}, []int{1, 0}},
			ans946{true},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans946, q.para946
				(validateStackSequences(p.one, p.two))
			}
		}
	}
}

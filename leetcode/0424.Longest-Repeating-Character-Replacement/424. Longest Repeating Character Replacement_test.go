package leetcode

import (
	"testing"
)

type question424 struct {
	para424
	ans424
}

// para 是参数
// one 代表第一个参数
type para424 struct {
	s string
	k int
}

// ans 是答案
// one 代表第一个答案
type ans424 struct {
	one int
}

func Benchmark_Problem424(b *testing.B) {

	qs := []question424{

		{
			para424{"AABABBA", 1},
			ans424{4},
		},

		{
			para424{"ABBB", 2},
			ans424{4},
		},

		{
			para424{"BAAA", 0},
			ans424{3},
		},

		{
			para424{"ABCDE", 1},
			ans424{2},
		},

		{
			para424{"BAAAB", 2},
			ans424{5},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans424, q.para424
				(characterReplacement(p.s, p.k))
			}
		}
	}
}

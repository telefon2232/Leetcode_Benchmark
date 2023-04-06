package leetcode

import (
	"testing"
)

type question60 struct {
	para60
	ans60
}

// para 是参数
// one 代表第一个参数
type para60 struct {
	n int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans60 struct {
	one string
}

func Benchmark_Problem60(b *testing.B) {

	qs := []question60{

		{
			para60{3, 3},
			ans60{"213"},
		},

		{
			para60{4, 9},
			ans60{"2314"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans60, q.para60
				(getPermutation(p.n, p.k))
			}
		}
	}
}

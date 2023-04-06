package leetcode

import (
	"testing"
)

type question132 struct {
	para132
	ans132
}

// para 是参数
// one 代表第一个参数
type para132 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans132 struct {
	one int
}

func Benchmark_Problem132(b *testing.B) {

	qs := []question132{

		{
			para132{"aab"},
			ans132{1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans132, q.para132
				(minCut(p.s))
			}
		}
	}
}

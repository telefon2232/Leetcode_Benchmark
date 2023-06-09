package leetcode

import (
	"testing"
)

type question756 struct {
	para756
	ans756
}

// para 是参数
// one 代表第一个参数
type para756 struct {
	b string
	a []string
}

// ans 是答案
// one 代表第一个答案
type ans756 struct {
	one bool
}

func Benchmark_Problem756(b *testing.B) {

	qs := []question756{

		{
			para756{"BCD", []string{"BCG", "CDE", "GEA", "FFF"}},
			ans756{true},
		},

		{
			para756{"AABA", []string{"AAA", "AAB", "ABA", "ABB", "BAC"}},
			ans756{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans756, q.para756
				(pyramidTransition(p.b, p.a))
			}
		}
	}
}

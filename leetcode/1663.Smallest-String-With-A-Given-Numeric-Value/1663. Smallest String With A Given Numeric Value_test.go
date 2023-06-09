package leetcode

import (
	"testing"
)

type question1663 struct {
	para1663
	ans1663
}

// para 是参数
// one 代表第一个参数
type para1663 struct {
	n int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans1663 struct {
	one string
}

func Benchmark_Problem1663(b *testing.B) {

	qs := []question1663{

		{
			para1663{3, 27},
			ans1663{"aay"},
		},

		{
			para1663{5, 73},
			ans1663{"aaszz"},
		},

		{
			para1663{24, 552},
			ans1663{"aaszz"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1663, q.para1663
				(getSmallestString(p.n, p.k))
			}
		}
	}
}

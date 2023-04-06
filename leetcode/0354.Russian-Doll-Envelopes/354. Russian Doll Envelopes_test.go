package leetcode

import (
	"testing"
)

type question354 struct {
	para354
	ans354
}

// para 是参数
// one 代表第一个参数
type para354 struct {
	envelopes [][]int
}

// ans 是答案
// one 代表第一个答案
type ans354 struct {
	one int
}

func Benchmark_Problem354(b *testing.B) {

	qs := []question354{

		{
			para354{[][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}},
			ans354{3},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans354, q.para354
				(maxEnvelopes(p.envelopes))
			}
		}
	}
}

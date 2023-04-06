package leetcode

import (
	"testing"
)

type question767 struct {
	para767
	ans767
}

// para 是参数
// one 代表第一个参数
type para767 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans767 struct {
	one string
}

func Benchmark_Problem767(b *testing.B) {

	qs := []question767{
		{
			para767{"aab"},
			ans767{"aba"},
		},

		{
			para767{"aaab"},
			ans767{""},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans767, q.para767
				(reorganizeString(p.one))
			}
		}
	}
}

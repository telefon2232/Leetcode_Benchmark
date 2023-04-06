package leetcode

import (
	"testing"
)

type question115 struct {
	para115
	ans115
}

// para 是参数
// one 代表第一个参数
type para115 struct {
	s string
	t string
}

// ans 是答案
// one 代表第一个答案
type ans115 struct {
	one int
}

func Benchmark_Problem115(b *testing.B) {

	qs := []question115{

		{
			para115{"rabbbit", "rabbit"},
			ans115{3},
		},

		{
			para115{"babgbag", "bag"},
			ans115{5},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans115, q.para115
				(numDistinct(p.s, p.t))
			}
		}
	}
}

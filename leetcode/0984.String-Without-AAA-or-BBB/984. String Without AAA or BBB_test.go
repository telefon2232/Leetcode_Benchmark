package leetcode

import (
	"testing"
)

type question984 struct {
	para984
	ans984
}

// para 是参数
// one 代表第一个参数
type para984 struct {
	a int
	b int
}

// ans 是答案
// one 代表第一个答案
type ans984 struct {
	one string
}

func Benchmark_Problem984(b *testing.B) {

	qs := []question984{

		{
			para984{1, 2},
			ans984{"abb"},
		},

		{
			para984{4, 1},
			ans984{"aabaa"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans984, q.para984
				(strWithout3a3b(p.a, p.b))
			}
		}
	}
}

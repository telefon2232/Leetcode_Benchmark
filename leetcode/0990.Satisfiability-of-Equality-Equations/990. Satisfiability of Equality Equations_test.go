package leetcode

import (
	"testing"
)

type question990 struct {
	para990
	ans990
}

// para 是参数
// one 代表第一个参数
type para990 struct {
	a []string
}

// ans 是答案
// one 代表第一个答案
type ans990 struct {
	one bool
}

func Benchmark_Problem990(b *testing.B) {

	qs := []question990{

		{
			para990{[]string{"a==b", "b!=a"}},
			ans990{false},
		},

		{
			para990{[]string{"b==a", "a==b"}},
			ans990{true},
		},

		{
			para990{[]string{"a==b", "b==c", "a==c"}},
			ans990{true},
		},

		{
			para990{[]string{"a==b", "b!=c", "c==a"}},
			ans990{false},
		},

		{
			para990{[]string{"c==c", "b==d", "x!=z"}},
			ans990{true},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans990, q.para990
				(equationsPossible(p.a))
			}
		}
	}
}

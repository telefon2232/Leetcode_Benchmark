package leetcode

import (
	"testing"
)

type question93 struct {
	para93
	ans93
}

// para 是参数
// one 代表第一个参数
type para93 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans93 struct {
	one []string
}

func Benchmark_Problem93(b *testing.B) {

	qs := []question93{

		{
			para93{"25525511135"},
			ans93{[]string{"255.255.11.135", "255.255.111.35"}},
		},

		{
			para93{"0000"},
			ans93{[]string{"0.0.0.0"}},
		},

		{
			para93{"010010"},
			ans93{[]string{"0.10.0.10", "0.100.1.0"}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans93, q.para93
				(restoreIPAddresses(p.s))
			}
		}
	}
}

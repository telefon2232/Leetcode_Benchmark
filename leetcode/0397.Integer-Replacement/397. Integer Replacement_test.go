package leetcode

import (
	"testing"
)

type question397 struct {
	para397
	ans397
}

// para 是参数
// one 代表第一个参数
type para397 struct {
	s int
}

// ans 是答案
// one 代表第一个答案
type ans397 struct {
	one int
}

func Benchmark_Problem397(b *testing.B) {

	qs := []question397{

		{
			para397{8},
			ans397{3},
		},

		{
			para397{7},
			ans397{4},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans397, q.para397
				(integerReplacement(p.s))
			}
		}
	}
}

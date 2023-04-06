package leetcode

import (
	"testing"
)

type question342 struct {
	para342
	ans342
}

// para 是参数
// one 代表第一个参数
type para342 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans342 struct {
	one bool
}

func Benchmark_Problem342(b *testing.B) {

	qs := []question342{

		{
			para342{16},
			ans342{true},
		},

		{
			para342{5},
			ans342{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans342, q.para342
				(isPowerOfFour(p.one))
			}
		}
	}
}

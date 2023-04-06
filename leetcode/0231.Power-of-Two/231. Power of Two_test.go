package leetcode

import (
	"testing"
)

type question231 struct {
	para231
	ans231
}

// para 是参数
// one 代表第一个参数
type para231 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans231 struct {
	one bool
}

func Benchmark_Problem231(b *testing.B) {

	qs := []question231{

		{
			para231{1},
			ans231{true},
		},

		{
			para231{16},
			ans231{true},
		},

		{
			para231{218},
			ans231{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans231, q.para231
				(isPowerOfTwo(p.one))
			}
		}
	}
}

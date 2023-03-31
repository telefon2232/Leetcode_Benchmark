package leetcode

import (
	"testing"
)

type question921 struct {
	para921
	ans921
}

// para 是参数
// one 代表第一个参数
type para921 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans921 struct {
	one int
}

func Benchmark_Problem921(b *testing.B) {

	qs := []question921{
		{
			para921{"())"},
			ans921{1},
		},

		{
			para921{"((("},
			ans921{3},
		},

		{
			para921{"()"},
			ans921{0},
		},

		{
			para921{"()))(("},
			ans921{4},
		},
	}

	for _, q := range qs {
		_, p := q.ans921, q.para921
		(minAddToMakeValid(p.one))
	}
}

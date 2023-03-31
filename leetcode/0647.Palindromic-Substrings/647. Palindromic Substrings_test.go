package leetcode

import (
	"testing"
)

type question647 struct {
	para647
	ans647
}

// para 是参数
// one 代表第一个参数
type para647 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans647 struct {
	one int
}

func Benchmark_Problem647(b *testing.B) {

	qs := []question647{

		{
			para647{"abc"},
			ans647{3},
		},

		{
			para647{"aaa"},
			ans647{6},
		},
	}

	for _, q := range qs {
		_, p := q.ans647, q.para647
		(countSubstrings(p.s))
	}
}

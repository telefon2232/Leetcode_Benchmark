package leetcode

import (
	"testing"
)

type question224 struct {
	para224
	ans224
}

// para 是参数
// one 代表第一个参数
type para224 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans224 struct {
	one int
}

func Benchmark_Problem224(b *testing.B) {

	qs := []question224{

		{
			para224{"1 + 1"},
			ans224{2},
		},
		{
			para224{" 2-1 + 2 "},
			ans224{3},
		},

		{
			para224{"(1+(4+5+2)-3)+(6+8)"},
			ans224{23},
		},

		{
			para224{"2-(5-6)"},
			ans224{3},
		},
	}

	for _, q := range qs {
		_, p := q.ans224, q.para224
		(calculate(p.one))
	}
}

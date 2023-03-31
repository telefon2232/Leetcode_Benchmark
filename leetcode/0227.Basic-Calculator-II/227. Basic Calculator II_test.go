package leetcode

import (
	"testing"
)

type question227 struct {
	para227
	ans227
}

// para 是参数
// one 代表第一个参数
type para227 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans227 struct {
	one int
}

func Benchmark_Problem227(b *testing.B) {

	qs := []question227{

		{
			para227{"3+2*2"},
			ans227{7},
		},

		{
			para227{"3/2"},
			ans227{1},
		},

		{
			para227{" 3+5 / 2 "},
			ans227{5},
		},

		{
			para227{"1 + 1"},
			ans227{2},
		},
		{
			para227{" 2-1 + 2 "},
			ans227{3},
		},

		{
			para227{"2-5/6"},
			ans227{2},
		},
	}

	for _, q := range qs {
		_, p := q.ans227, q.para227
		(calculate(p.one))
	}
}

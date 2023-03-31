package leetcode

import (
	"testing"
)

type question202 struct {
	para202
	ans202
}

// para 是参数
// one 代表第一个参数
type para202 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans202 struct {
	one bool
}

func Benchmark_Problem202(b *testing.B) {

	qs := []question202{

		{
			para202{202},
			ans202{false},
		},

		{
			para202{19},
			ans202{true},
		},

		{
			para202{2},
			ans202{false},
		},

		{
			para202{3},
			ans202{false},
		},
	}

	for _, q := range qs {
		_, p := q.ans202, q.para202
		(isHappy(p.one))
	}
}

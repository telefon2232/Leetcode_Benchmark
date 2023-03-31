package leetcode

import (
	"testing"
)

type question28 struct {
	para28
	ans28
}

// para 是参数
// one 代表第一个参数
type para28 struct {
	s string
	p string
}

// ans 是答案
// one 代表第一个答案
type ans28 struct {
	one int
}

func Benchmark_Problem28(b *testing.B) {

	qs := []question28{

		{
			para28{"abab", "ab"},
			ans28{0},
		},

		{
			para28{"hello", "ll"},
			ans28{2},
		},

		{
			para28{"", "abc"},
			ans28{0},
		},

		{
			para28{"abacbabc", "abc"},
			ans28{5},
		},

		{
			para28{"abacbabc", "abcd"},
			ans28{-1},
		},

		{
			para28{"abacbabc", ""},
			ans28{0},
		},
	}

	for _, q := range qs {
		_, p := q.ans28, q.para28
		(strStr(p.s, p.p))
	}
}

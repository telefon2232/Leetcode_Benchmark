package leetcode

import (
	"testing"
)

type question541 struct {
	para541
	ans541
}

// para 是参数
// one 代表第一个参数
type para541 struct {
	s string
	k int
}

// ans 是答案
// one 代表第一个答案
type ans541 struct {
	one string
}

func Benchmark_Problem541(b *testing.B) {

	qs := []question541{

		{
			para541{"abcdefg", 2},
			ans541{"bacdfeg"},
		},

		{
			para541{"abcdefg", 5},
			ans541{"edcbafg"},
		},

		{
			para541{"abcd", 4},
			ans541{"dcba"},
		},

		{
			para541{"", 100},
			ans541{""},
		},
	}

	for _, q := range qs {
		_, p := q.ans541, q.para541
		(reverseStr(p.s, p.k))
	}
}

package leetcode

import (
	"testing"
)

type question1003 struct {
	para1003
	ans1003
}

// para 是参数
// one 代表第一个参数
type para1003 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans1003 struct {
	one bool
}

func Benchmark_Problem1003(b *testing.B) {

	qs := []question1003{

		{
			para1003{"aabcbc"},
			ans1003{true},
		},

		{
			para1003{"abcabcababcc"},
			ans1003{true},
		},

		{
			para1003{"abccba"},
			ans1003{false},
		},

		{
			para1003{"cababc"},
			ans1003{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans1003, q.para1003
		(isValid1003(p.s))
	}
}}}

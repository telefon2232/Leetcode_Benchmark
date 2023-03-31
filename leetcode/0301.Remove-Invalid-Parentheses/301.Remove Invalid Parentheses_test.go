package leetcode

import (
	"testing"
)

type question301 struct {
	para301
	ans301
}

// s 是参数
type para301 struct {
	s string
}

// ans 是答案
type ans301 struct {
	ans []string
}

func Benchmark_Problem301(b *testing.B) {

	qs := []question301{

		{
			para301{"()())()"},
			ans301{[]string{"(())()", "()()()"}},
		},

		{
			para301{"(a)())()"},
			ans301{[]string{"(a())()", "(a)()()"}},
		},

		{
			para301{")("},
			ans301{[]string{""}},
		},
	}

	for _, q := range qs {
		_, p := q.ans301, q.para301
		(removeInvalidParentheses(p.s))
	}
}

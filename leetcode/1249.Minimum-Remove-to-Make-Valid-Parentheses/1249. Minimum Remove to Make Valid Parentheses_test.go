package leetcode

import (
	"testing"
)

type question1249 struct {
	para1249
	ans1249
}

// para 是参数
// one 代表第一个参数
type para1249 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans1249 struct {
	one string
}

func Benchmark_Problem1249(b *testing.B) {

	qs := []question1249{

		{
			para1249{"lee(t(c)o)de)"},
			ans1249{"lee(t(c)o)de"},
		},

		{
			para1249{"a)b(c)d"},
			ans1249{"ab(c)d"},
		},

		{
			para1249{"))(("},
			ans1249{""},
		},

		{
			para1249{"(a(b(c)d)"},
			ans1249{"a(b(c)d)"},
		},
	}

	for _, q := range qs {
		_, p := q.ans1249, q.para1249
		(minRemoveToMakeValid(p.s))
	}
}

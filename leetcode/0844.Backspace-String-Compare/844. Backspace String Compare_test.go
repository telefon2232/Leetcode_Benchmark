package leetcode

import (
	"testing"
)

type question844 struct {
	para844
	ans844
}

// para 是参数
// one 代表第一个参数
type para844 struct {
	s string
	t string
}

// ans 是答案
// one 代表第一个答案
type ans844 struct {
	one bool
}

func Benchmark_Problem844(b *testing.B) {

	qs := []question844{

		{
			para844{"ab#c", "ad#c"},
			ans844{true},
		},

		{
			para844{"ab##", "c#d#"},
			ans844{true},
		},

		{
			para844{"a##c", "#a#c"},
			ans844{true},
		},

		{
			para844{"a#c", "b"},
			ans844{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans844, q.para844
		(backspaceCompare(p.s, p.t))
	}
}}}

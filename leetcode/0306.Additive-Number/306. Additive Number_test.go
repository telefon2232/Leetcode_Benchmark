package leetcode

import (
	"testing"
)

type question306 struct {
	para306
	ans306
}

// para 是参数
// one 代表第一个参数
type para306 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans306 struct {
	one bool
}

func Benchmark_Problem306(b *testing.B) {

	qs := []question306{

		{
			para306{"112358"},
			ans306{true},
		},

		{
			para306{"199100199"},
			ans306{true},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans306, q.para306
		(isAdditiveNumber(p.one))
	}
}}}

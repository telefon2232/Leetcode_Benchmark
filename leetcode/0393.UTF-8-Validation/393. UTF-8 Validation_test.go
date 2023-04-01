package leetcode

import (
	"testing"
)

type question393 struct {
	para393
	ans393
}

// para 是参数
// one 代表第一个参数
type para393 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans393 struct {
	one bool
}

func Benchmark_Problem393(b *testing.B) {

	qs := []question393{

		{
			para393{[]int{197, 130, 1}},
			ans393{true},
		},

		{
			para393{[]int{235, 140, 4}},
			ans393{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans393, q.para393
		(validUtf8(p.one))
	}
}}}

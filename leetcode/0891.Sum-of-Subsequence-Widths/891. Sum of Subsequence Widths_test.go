package leetcode

import (
	"testing"
)

type question891 struct {
	para891
	ans891
}

// para 是参数
// one 代表第一个参数
type para891 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans891 struct {
	one int
}

func Benchmark_Problem891(b *testing.B) {

	qs := []question891{

		{
			para891{[]int{2, 1, 3}},
			ans891{6},
		},

		{
			para891{[]int{3, 7, 2, 3}},
			ans891{35},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans891, q.para891
		(sumSubseqWidths(p.one))
	}
}}}

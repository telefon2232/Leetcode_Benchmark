package leetcode

import (
	"testing"
)

type question201 struct {
	para201
	ans201
}

// para 是参数
// one 代表第一个参数
type para201 struct {
	m int
	n int
}

// ans 是答案
// one 代表第一个答案
type ans201 struct {
	one int
}

func Benchmark_Problem201(b *testing.B) {

	qs := []question201{

		{
			para201{5, 7},
			ans201{4},
		},

		{
			para201{0, 1},
			ans201{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans201, q.para201
		(rangeBitwiseAnd(p.m, p.n))
	}
}}}

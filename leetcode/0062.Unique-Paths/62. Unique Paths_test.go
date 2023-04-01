package leetcode

import (
	"testing"
)

type question62 struct {
	para62
	ans62
}

// para 是参数
// one 代表第一个参数
type para62 struct {
	m int
	n int
}

// ans 是答案
// one 代表第一个答案
type ans62 struct {
	one int
}

func Benchmark_Problem62(b *testing.B) {

	qs := []question62{

		{
			para62{3, 2},
			ans62{3},
		},

		{
			para62{7, 3},
			ans62{28},
		},

		{
			para62{1, 2},
			ans62{1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans62, q.para62
		(uniquePaths(p.m, p.n))
	}
}}}

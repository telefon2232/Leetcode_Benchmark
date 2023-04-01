package leetcode

import (
	"testing"
)

type question576 struct {
	para576
	ans576
}

// para 是参数
// one 代表第一个参数
type para576 struct {
	m           int
	n           int
	maxMove     int
	startRow    int
	startColumn int
}

// ans 是答案
// one 代表第一个答案
type ans576 struct {
	one int
}

func Benchmark_Problem576(b *testing.B) {

	qs := []question576{

		{
			para576{2, 2, 2, 0, 0},
			ans576{6},
		},

		{
			para576{1, 3, 3, 0, 1},
			ans576{12},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans576, q.para576
		(findPaths(p.m, p.n, p.maxMove, p.startRow, p.startColumn))
	}
}}}

package leetcode

import (
	"testing"
)

type question785 struct {
	para785
	ans785
}

// para 是参数
// one 代表第一个参数
type para785 struct {
	graph [][]int
}

// ans 是答案
// one 代表第一个答案
type ans785 struct {
	one bool
}

func Benchmark_Problem785(b *testing.B) {

	qs := []question785{

		{
			para785{[][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}},
			ans785{true},
		},

		{
			para785{[][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}},
			ans785{false},
		},

		{
			para785{[][]int{{1, 2, 3}, {0, 2, 3}, {0, 1, 3}, {0, 1, 2}}},
			ans785{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans785, q.para785
		(isBipartite(p.graph))
	}
}}}

package leetcode

import (
	"testing"
)

type question685 struct {
	para685
	ans685
}

// para 是参数
// one 代表第一个参数
type para685 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans685 struct {
	one []int
}

func Benchmark_Problem685(b *testing.B) {

	qs := []question685{

		{
			para685{[][]int{{3, 5}, {1, 3}, {2, 1}, {5, 4}, {2, 3}}},
			ans685{[]int{2, 3}},
		},

		{
			para685{[][]int{{4, 2}, {1, 5}, {5, 2}, {5, 3}, {2, 4}}},
			ans685{[]int{4, 2}},
		},

		{
			para685{[][]int{{1, 2}, {1, 3}, {2, 3}}},
			ans685{[]int{2, 3}},
		},

		{
			para685{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}},
			ans685{[]int{1, 4}},
		},

		{
			para685{[][]int{{2, 1}, {3, 1}, {4, 2}, {1, 4}}},
			ans685{[]int{2, 1}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans685, q.para685
		(findRedundantDirectedConnection(p.one))
	}
}}}

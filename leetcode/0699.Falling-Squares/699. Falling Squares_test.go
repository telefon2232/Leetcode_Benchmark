package leetcode

import (
	"testing"
)

type question699 struct {
	para699
	ans699
}

// para 是参数
// one 代表第一个参数
type para699 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans699 struct {
	one []int
}

func Benchmark_Problem699(b *testing.B) {

	qs := []question699{

		{
			para699{[][]int{{6, 1}, {9, 2}, {2, 4}}},
			ans699{[]int{1, 2, 4}},
		},

		{
			para699{[][]int{{1, 2}, {1, 3}}},
			ans699{[]int{2, 5}},
		},

		{
			para699{[][]int{{1, 2}, {2, 3}, {6, 1}}},
			ans699{[]int{2, 5, 5}},
		},

		{
			para699{[][]int{{100, 100}, {200, 100}}},
			ans699{[]int{100, 100}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans699, q.para699
		(fallingSquares(p.one))
	}
}}}

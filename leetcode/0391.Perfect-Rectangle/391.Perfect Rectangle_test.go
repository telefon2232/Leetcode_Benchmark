package leetcode

import (
	"testing"
)

type question391 struct {
	para391
	ans391
}

// para 是参数
type para391 struct {
	rectangles [][]int
}

// ans 是答案
type ans391 struct {
	ans bool
}

func Benchmark_Problem391(b *testing.B) {

	qs := []question391{

		{
			para391{[][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {3, 2, 4, 4}, {1, 3, 2, 4}, {2, 3, 3, 4}}},
			ans391{true},
		},

		{
			para391{[][]int{{1, 1, 2, 3}, {1, 3, 2, 4}, {3, 1, 4, 2}, {3, 2, 4, 4}}},
			ans391{false},
		},

		{
			para391{[][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {1, 3, 2, 4}, {3, 2, 4, 4}}},
			ans391{false},
		},

		{
			para391{[][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {1, 3, 2, 4}, {2, 2, 4, 4}}},
			ans391{false},
		},
	}

	for _, q := range qs {
		_, p := q.ans391, q.para391
		(isRectangleCover(p.rectangles))
	}
}

package leetcode

import (
	"testing"
)

type question832 struct {
	para832
	ans832
}

// para 是参数
// one 代表第一个参数
type para832 struct {
	A [][]int
}

// ans 是答案
// one 代表第一个答案
type ans832 struct {
	one [][]int
}

func Benchmark_Problem832(b *testing.B) {

	qs := []question832{

		{
			para832{[][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}},
			ans832{[][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
		},

		{
			para832{[][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}},
			ans832{[][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}}},
		},

		{
			para832{[][]int{{1, 1, 1}, {1, 1, 1}, {0, 0, 0}}},
			ans832{[][]int{{0, 0, 0}, {0, 0, 0}, {1, 1, 1}}},
		},
	}

	for _, q := range qs {
		_, p := q.ans832, q.para832
		(flipAndInvertImage(p.A))
	}
}

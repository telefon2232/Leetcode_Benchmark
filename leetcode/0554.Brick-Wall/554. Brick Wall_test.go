package leetcode

import (
	"testing"
)

type question554 struct {
	para554
	ans554
}

// para 是参数
// one 代表第一个参数
type para554 struct {
	wall [][]int
}

// ans 是答案
// one 代表第一个答案
type ans554 struct {
	one int
}

func Benchmark_Problem554(b *testing.B) {

	qs := []question554{

		{
			para554{[][]int{{1, 2, 2, 1}, {3, 1, 2}, {1, 3, 2}, {2, 4}, {3, 1, 2}, {1, 3, 1, 1}}},
			ans554{2},
		},

		{
			para554{[][]int{{1}, {1}, {1}}},
			ans554{3},
		},

		{
			para554{[][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}},
			ans554{1},
		},

		{
			para554{[][]int{{1, 1, 0}, {1, 1, 1}, {0, 1, 1}}},
			ans554{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans554, q.para554
				(leastBricks(p.wall))
			}
		}
	}
}

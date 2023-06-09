package leetcode

import (
	"testing"
)

type question947 struct {
	para947
	ans947
}

// para 是参数
// one 代表第一个参数
type para947 struct {
	stones [][]int
}

// ans 是答案
// one 代表第一个答案
type ans947 struct {
	one int
}

func Benchmark_Problem947(b *testing.B) {

	qs := []question947{
		{
			para947{[][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}},
			ans947{5},
		},

		{
			para947{[][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}},
			ans947{3},
		},

		{
			para947{[][]int{{0, 0}}},
			ans947{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans947, q.para947
				(removeStones(p.stones))
			}
		}
	}
}

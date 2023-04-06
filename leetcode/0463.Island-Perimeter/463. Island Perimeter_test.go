package leetcode

import (
	"testing"
)

type question463 struct {
	para463
	ans463
}

// para 是参数
// one 代表第一个参数
type para463 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans463 struct {
	one int
}

func Benchmark_Problem463(b *testing.B) {

	qs := []question463{

		{
			para463{[][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}},
			ans463{16},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans463, q.para463
				(islandPerimeter(p.one))
			}
		}
	}
}

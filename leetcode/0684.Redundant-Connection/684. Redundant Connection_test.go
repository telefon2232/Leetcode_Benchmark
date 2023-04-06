package leetcode

import (
	"testing"
)

type question684 struct {
	para684
	ans684
}

// para 是参数
// one 代表第一个参数
type para684 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans684 struct {
	one []int
}

func Benchmark_Problem684(b *testing.B) {

	qs := []question684{

		{
			para684{[][]int{{1, 2}, {1, 3}, {2, 3}}},
			ans684{[]int{2, 3}},
		},

		{
			para684{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}},
			ans684{[]int{1, 4}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans684, q.para684
				(findRedundantConnection(p.one))
			}
		}
	}
}

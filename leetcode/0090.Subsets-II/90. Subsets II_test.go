package leetcode

import (
	"testing"
)

type question90 struct {
	para90
	ans90
}

// para 是参数
// one 代表第一个参数
type para90 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans90 struct {
	one [][]int
}

func Benchmark_Problem90(b *testing.B) {

	qs := []question90{

		{
			para90{[]int{}},
			ans90{[][]int{{}}},
		},

		{
			para90{[]int{1, 2, 2}},
			ans90{[][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans90, q.para90
				(subsetsWithDup(p.one))
			}
		}
	}
}

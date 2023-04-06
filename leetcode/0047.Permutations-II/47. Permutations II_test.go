package leetcode

import (
	"testing"
)

type question47 struct {
	para47
	ans47
}

// para 是参数
// one 代表第一个参数
type para47 struct {
	s []int
}

// ans 是答案
// one 代表第一个答案
type ans47 struct {
	one [][]int
}

func Benchmark_Problem47(b *testing.B) {

	qs := []question47{

		{
			para47{[]int{1, 1, 2}},
			ans47{[][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
		},

		{
			para47{[]int{1, 2, 2}},
			ans47{[][]int{{1, 2, 2}, {2, 2, 1}, {2, 1, 2}}},
		},

		{
			para47{[]int{2, 2, 2}},
			ans47{[][]int{{2, 2, 2}}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans47, q.para47
				(permuteUnique(p.s))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question96 struct {
	para96
	ans96
}

// para 是参数
// one 代表第一个参数
type para96 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans96 struct {
	one int
}

func Benchmark_Problem96(b *testing.B) {

	qs := []question96{

		{
			para96{1},
			ans96{1},
		},

		{
			para96{3},
			ans96{5},
		},

		{
			para96{4},
			ans96{14},
		},

		{
			para96{5},
			ans96{42},
		},

		{
			para96{6},
			ans96{132},
		},

		{
			para96{7},
			ans96{429},
		},

		{
			para96{8},
			ans96{1430},
		},

		{
			para96{9},
			ans96{4862},
		},
		{
			para96{10},
			ans96{16796},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans96, q.para96
				(numTrees(p.one))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question830 struct {
	para830
	ans830
}

// para 是参数
// one 代表第一个参数
type para830 struct {
	S string
}

// ans 是答案
// one 代表第一个答案
type ans830 struct {
	one [][]int
}

func Benchmark_Problem830(b *testing.B) {

	qs := []question830{

		{
			para830{"abbxxxxzzy"},
			ans830{[][]int{{3, 6}}},
		},

		{
			para830{"abc"},
			ans830{[][]int{{}}},
		},

		{
			para830{"abcdddeeeeaabbbcd"},
			ans830{[][]int{{3, 5}, {6, 9}, {12, 14}}},
		},

		{
			para830{"aba"},
			ans830{[][]int{{}}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans830, q.para830
				(largeGroupPositions(p.S))
			}
		}
	}
}

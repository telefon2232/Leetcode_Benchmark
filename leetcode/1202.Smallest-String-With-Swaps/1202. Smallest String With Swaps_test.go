package leetcode

import (
	"testing"
)

type question1202 struct {
	para1202
	ans1202
}

// para 是参数
// one 代表第一个参数
type para1202 struct {
	s     string
	pairs [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1202 struct {
	one string
}

func Benchmark_Problem1202(b *testing.B) {

	qs := []question1202{

		{
			para1202{"dcab", [][]int{{0, 3}, {1, 2}}},
			ans1202{"bacd"},
		},

		{
			para1202{"dcab", [][]int{{0, 3}, {1, 2}, {0, 2}}},
			ans1202{"abcd"},
		},

		{
			para1202{"cba", [][]int{{0, 1}, {1, 2}}},
			ans1202{"abc"},
		},
	}

	for _, q := range qs {
		_, p := q.ans1202, q.para1202
		(smallestStringWithSwaps(p.s, p.pairs))
	}
}

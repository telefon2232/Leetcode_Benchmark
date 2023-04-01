package leetcode

import (
	"testing"
)

type question216 struct {
	para216
	ans216
}

// para 是参数
// one 代表第一个参数
type para216 struct {
	n int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans216 struct {
	one [][]int
}

func Benchmark_Problem216(b *testing.B) {

	qs := []question216{

		{
			para216{3, 7},
			ans216{[][]int{{1, 2, 4}}},
		},
		{
			para216{3, 9},
			ans216{[][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans216, q.para216
		(combinationSum3(p.n, p.k))
	}
}}}

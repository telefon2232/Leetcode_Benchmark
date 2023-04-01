package leetcode

import (
	"testing"
)

type question1329 struct {
	para1329
	ans1329
}

// para 是参数
// one 代表第一个参数
type para1329 struct {
	mat [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1329 struct {
	one [][]int
}

func Benchmark_Problem1329(b *testing.B) {

	qs := []question1329{

		{
			para1329{[][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}}},
			ans1329{[][]int{{1, 1, 1, 1}, {1, 2, 2, 2}, {1, 2, 3, 3}}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans1329, q.para1329
		(diagonalSort(p.mat))
	}
}}}

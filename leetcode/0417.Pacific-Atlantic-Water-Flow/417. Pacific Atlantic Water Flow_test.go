package leetcode

import (
	"testing"
)

type question417 struct {
	para417
	ans417
}

// para 是参数
// one 代表第一个参数
type para417 struct {
	matrix [][]int
}

// ans 是答案
// one 代表第一个答案
type ans417 struct {
	one [][]int
}

func Benchmark_Problem417(b *testing.B) {

	qs := []question417{

		{
			para417{[][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			}},
			ans417{[][]int{
				{0, 4},
				{1, 3},
				{1, 4},
				{2, 2},
				{3, 0},
				{3, 1},
				{4, 0},
			}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans417, q.para417
		(pacificAtlantic(p.matrix))
	}
}}}

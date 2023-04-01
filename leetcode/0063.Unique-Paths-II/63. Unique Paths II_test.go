package leetcode

import (
	"testing"
)

type question63 struct {
	para63
	ans63
}

// para 是参数
// one 代表第一个参数
type para63 struct {
	og [][]int
}

// ans 是答案
// one 代表第一个答案
type ans63 struct {
	one int
}

func Benchmark_Problem63(b *testing.B) {

	qs := []question63{

		{
			para63{[][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			}},
			ans63{2},
		},

		{
			para63{[][]int{
				{0, 0},
				{1, 1},
				{0, 0},
			}},
			ans63{0},
		},

		{
			para63{[][]int{
				{0, 1, 0, 0},
				{1, 0, 0, 0},
				{0, 0, 0, 0},
			}},
			ans63{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans63, q.para63
		(uniquePathsWithObstacles(p.og))
	}
}}}

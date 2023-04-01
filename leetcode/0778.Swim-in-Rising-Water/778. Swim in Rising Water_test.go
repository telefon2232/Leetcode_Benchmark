package leetcode

import (
	"testing"
)

type question778 struct {
	para778
	ans778
}

// para 是参数
// one 代表第一个参数
type para778 struct {
	grid [][]int
}

// ans 是答案
// one 代表第一个答案
type ans778 struct {
	one int
}

func Benchmark_Problem778(b *testing.B) {

	qs := []question778{
		{
			para778{[][]int{{0, 2}, {1, 3}}},
			ans778{3},
		},

		{
			para778{[][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}}},
			ans778{16},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans778, q.para778
		(swimInWater(p.grid))
	}
}}}

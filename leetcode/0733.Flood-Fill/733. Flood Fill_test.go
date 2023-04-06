package leetcode

import (
	"testing"
)

type question733 struct {
	para733
	ans733
}

// para 是参数
// one 代表第一个参数
type para733 struct {
	one [][]int
	sr  int
	sc  int
	c   int
}

// ans 是答案
// one 代表第一个答案
type ans733 struct {
	one [][]int
}

func Benchmark_Problem733(b *testing.B) {

	qs := []question733{

		{
			para733{[][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
			}, 1, 1, 2},
			ans733{[][]int{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans733, q.para733
				(floodFill(p.one, p.sr, p.sc, p.c))
			}
		}
	}
}

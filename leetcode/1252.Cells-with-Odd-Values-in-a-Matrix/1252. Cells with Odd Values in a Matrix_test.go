package leetcode

import (
	"testing"
)

type question1252 struct {
	para1252
	ans1252
}

// para 是参数
// one 代表第一个参数
type para1252 struct {
	n       int
	m       int
	indices [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1252 struct {
	one int
}

func Benchmark_Problem1252(b *testing.B) {

	qs := []question1252{

		{
			para1252{2, 3, [][]int{{0, 1}, {1, 1}}},
			ans1252{6},
		},

		{
			para1252{2, 2, [][]int{{1, 1}, {0, 0}}},
			ans1252{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1252, q.para1252
				(oddCells(p.n, p.m, p.indices))
			}
		}
	}
}

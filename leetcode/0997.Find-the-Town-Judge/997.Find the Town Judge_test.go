package leetcode

import (
	"testing"
)

type question997 struct {
	para997
	ans997
}

// para 是参数
type para997 struct {
	n     int
	trust [][]int
}

// ans 是答案
type ans997 struct {
	ans int
}

func Benchmark_Problem997(b *testing.B) {

	qs := []question997{

		{
			para997{2, [][]int{{1, 2}}},
			ans997{2},
		},

		{
			para997{3, [][]int{{1, 3}, {2, 3}}},
			ans997{3},
		},

		{
			para997{3, [][]int{{1, 3}, {2, 3}, {3, 1}}},
			ans997{-1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans997, q.para997
				(findJudge(p.n, p.trust))
			}
		}
	}
}

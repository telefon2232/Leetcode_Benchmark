package leetcode

import (
	"testing"
)

type question1486 struct {
	para1486
	ans1486
}

// para 是参数
// one 代表第一个参数
type para1486 struct {
	n     int
	start int
}

// ans 是答案
// one 代表第一个答案
type ans1486 struct {
	one int
}

func Benchmark_Problem1486(b *testing.B) {

	qs := []question1486{

		{
			para1486{5, 0},
			ans1486{8},
		},

		{
			para1486{4, 3},
			ans1486{8},
		},

		{
			para1486{1, 7},
			ans1486{7},
		},

		{
			para1486{10, 5},
			ans1486{2},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1486, q.para1486
				(xorOperation(p.n, p.start))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question1659 struct {
	para1659
	ans1659
}

// para 是参数
// one 代表第一个参数
type para1659 struct {
	m               int
	n               int
	introvertsCount int
	extrovertsCount int
}

// ans 是答案
// one 代表第一个答案
type ans1659 struct {
	one int
}

func Benchmark_Problem1659(b *testing.B) {

	qs := []question1659{

		{
			para1659{2, 3, 1, 2},
			ans1659{240},
		},

		{
			para1659{3, 1, 2, 1},
			ans1659{260},
		},

		{
			para1659{2, 2, 4, 0},
			ans1659{240},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1659, q.para1659
				(getMaxGridHappiness(p.m, p.n, p.introvertsCount, p.extrovertsCount))
			}
		}
	}
}

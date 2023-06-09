package leetcode

import (
	"testing"
)

type question1641 struct {
	para1641
	ans1641
}

// para 是参数
// one 代表第一个参数
type para1641 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans1641 struct {
	one int
}

func Benchmark_Problem1641(b *testing.B) {

	qs := []question1641{

		{
			para1641{1},
			ans1641{5},
		},

		{
			para1641{2},
			ans1641{15},
		},

		{
			para1641{33},
			ans1641{66045},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1641, q.para1641
				(countVowelStrings(p.n))
			}
		}
	}
}

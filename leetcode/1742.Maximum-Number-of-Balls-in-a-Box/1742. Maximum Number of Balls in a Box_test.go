package leetcode

import (
	"testing"
)

type question1742 struct {
	para1742
	ans1742
}

// para 是参数
// one 代表第一个参数
type para1742 struct {
	lowLimit  int
	highLimit int
}

// ans 是答案
// one 代表第一个答案
type ans1742 struct {
	one int
}

func Benchmark_Problem1742(b *testing.B) {

	qs := []question1742{

		{
			para1742{1, 10},
			ans1742{2},
		},

		{
			para1742{5, 15},
			ans1742{2},
		},

		{
			para1742{19, 28},
			ans1742{2},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1742, q.para1742
				(countBalls(p.lowLimit, p.highLimit))
			}
		}
	}
}

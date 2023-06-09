package leetcode

import (
	"testing"
)

type question1646 struct {
	para1646
	ans1646
}

// para 是参数
// one 代表第一个参数
type para1646 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans1646 struct {
	one int
}

func Benchmark_Problem1646(b *testing.B) {

	qs := []question1646{

		{
			para1646{7},
			ans1646{3},
		},

		{
			para1646{2},
			ans1646{1},
		},

		{
			para1646{3},
			ans1646{2},
		},

		{
			para1646{0},
			ans1646{0},
		},

		{
			para1646{1},
			ans1646{1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1646, q.para1646
				(getMaximumGenerated(p.n))
			}
		}
	}
}

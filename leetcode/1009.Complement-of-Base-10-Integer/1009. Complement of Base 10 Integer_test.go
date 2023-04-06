package leetcode

import (
	"testing"
)

type question1009 struct {
	para1009
	ans1009
}

// para 是参数
// one 代表第一个参数
type para1009 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans1009 struct {
	one int
}

func Benchmark_Problem1009(b *testing.B) {

	qs := []question1009{

		{
			para1009{5},
			ans1009{2},
		},

		{
			para1009{7},
			ans1009{0},
		},

		{
			para1009{10},
			ans1009{5},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1009, q.para1009
				(bitwiseComplement(p.n))
			}
		}
	}
}

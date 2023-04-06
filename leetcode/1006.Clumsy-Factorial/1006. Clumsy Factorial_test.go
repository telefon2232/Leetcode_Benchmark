package leetcode

import (
	"testing"
)

type question1006 struct {
	para1006
	ans1006
}

// para 是参数
// one 代表第一个参数
type para1006 struct {
	N int
}

// ans 是答案
// one 代表第一个答案
type ans1006 struct {
	one int
}

func Benchmark_Problem1006(b *testing.B) {

	qs := []question1006{

		{
			para1006{4},
			ans1006{7},
		},

		{
			para1006{10},
			ans1006{12},
		},

		{
			para1006{100},
			ans1006{101},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1006, q.para1006
				(clumsy(p.N))
			}
		}
	}
}

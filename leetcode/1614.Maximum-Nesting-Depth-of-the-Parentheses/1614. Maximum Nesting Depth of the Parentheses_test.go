package leetcode

import (
	"testing"
)

type question1614 struct {
	para1614
	ans1614
}

// para 是参数
// one 代表第一个参数
type para1614 struct {
	p string
}

// ans 是答案
// one 代表第一个答案
type ans1614 struct {
	one int
}

func Benchmark_Problem1614(b *testing.B) {

	qs := []question1614{

		{
			para1614{"(1+(2*3)+((8)/4))+1"},
			ans1614{3},
		},

		{
			para1614{"(1)+((2))+(((3)))"},
			ans1614{3},
		},

		{
			para1614{"1+(2*3)/(2-1)"},
			ans1614{1},
		},

		{
			para1614{"1"},
			ans1614{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1614, q.para1614
				(maxDepth(p.p))
			}
		}
	}
}

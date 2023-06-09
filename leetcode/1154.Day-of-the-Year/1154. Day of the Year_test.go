package leetcode

import (
	"testing"
)

type question1154 struct {
	para1154
	ans1154
}

// para 是参数
// one 代表第一个参数
type para1154 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans1154 struct {
	one int
}

func Benchmark_Problem1154(b *testing.B) {

	qs := []question1154{

		{
			para1154{"2019-01-09"},
			ans1154{9},
		},

		{
			para1154{"2019-02-10"},
			ans1154{41},
		},

		{
			para1154{"2003-03-01"},
			ans1154{60},
		},

		{
			para1154{"2004-03-01"},
			ans1154{61},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1154, q.para1154
				(dayOfYear(p.one))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question1108 struct {
	para1108
	ans1108
}

// para 是参数
// one 代表第一个参数
type para1108 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans1108 struct {
	one string
}

func Benchmark_Problem1108(b *testing.B) {

	qs := []question1108{

		{
			para1108{"1.1.1.1"},
			ans1108{"1[.]1[.]1[.]1"},
		},

		{
			para1108{"255.100.50.0"},
			ans1108{"255[.]100[.]50[.]0"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1108, q.para1108
				(defangIPaddr(p.one))
			}
		}
	}
}

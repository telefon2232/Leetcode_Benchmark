package leetcode

import (
	"testing"
)

type question1185 struct {
	para1185
	ans1185
}

// para 是参数
// one 代表第一个参数
type para1185 struct {
	day   int
	month int
	year  int
}

// ans 是答案
// one 代表第一个答案
type ans1185 struct {
	one string
}

func Benchmark_Problem1185(b *testing.B) {

	qs := []question1185{

		{
			para1185{31, 8, 2019},
			ans1185{"Saturday"},
		},

		{
			para1185{18, 7, 1999},
			ans1185{"Sunday"},
		},

		{
			para1185{15, 8, 1993},
			ans1185{"Sunday"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans1185, q.para1185
		(dayOfTheWeek(p.day, p.month, p.year))
	}
}}}

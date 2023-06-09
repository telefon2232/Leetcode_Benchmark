package leetcode

import (
	"testing"
)

type question1704 struct {
	para1704
	ans1704
}

// para 是参数
// one 代表第一个参数
type para1704 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans1704 struct {
	one bool
}

func Benchmark_Problem1704(b *testing.B) {

	qs := []question1704{

		{
			para1704{"book"},
			ans1704{true},
		},

		{
			para1704{"textbook"},
			ans1704{false},
		},

		{
			para1704{"MerryChristmas"},
			ans1704{false},
		},

		{
			para1704{"AbCdEfGh"},
			ans1704{true},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1704, q.para1704
				(halvesAreAlike(p.s))
			}
		}
	}
}

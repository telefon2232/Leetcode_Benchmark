package leetcode

import (
	"testing"
)

type question1653 struct {
	para1653
	ans1653
}

// para 是参数
// one 代表第一个参数
type para1653 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans1653 struct {
	one int
}

func Benchmark_Problem1653(b *testing.B) {

	qs := []question1653{

		{
			para1653{"aababbab"},
			ans1653{2},
		},

		{
			para1653{"bbaaaaabb"},
			ans1653{2},
		},

		{
			para1653{"b"},
			ans1653{0},
		},

		{
			para1653{"ababaaaabbbbbaaababbbbbbaaabbaababbabbbbaabbbbaabbabbabaabbbababaa"},
			ans1653{25},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1653, q.para1653
				(minimumDeletions(p.s))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question1170 struct {
	para1170
	ans1170
}

// para 是参数
// one 代表第一个参数
type para1170 struct {
	queries []string
	words   []string
}

// ans 是答案
// one 代表第一个答案
type ans1170 struct {
	one []int
}

func Benchmark_Problem1170(b *testing.B) {

	qs := []question1170{

		{
			para1170{[]string{"cbd"}, []string{"zaaaz"}},
			ans1170{[]int{1}},
		},

		{
			para1170{[]string{"bbb", "cc"}, []string{"a", "aa", "aaa", "aaaa"}},
			ans1170{[]int{1, 2}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1170, q.para1170
				(numSmallerByFrequency(p.queries, p.words))
			}
		}
	}
}

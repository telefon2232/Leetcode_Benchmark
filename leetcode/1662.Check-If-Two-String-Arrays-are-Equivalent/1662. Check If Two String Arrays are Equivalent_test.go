package leetcode

import (
	"testing"
)

type question1662 struct {
	para1662
	ans1662
}

// para 是参数
// one 代表第一个参数
type para1662 struct {
	word1 []string
	word2 []string
}

// ans 是答案
// one 代表第一个答案
type ans1662 struct {
	one bool
}

func Benchmark_Problem1662(b *testing.B) {

	qs := []question1662{

		{
			para1662{[]string{"ab", "c"}, []string{"a", "bc"}},
			ans1662{true},
		},

		{
			para1662{[]string{"a", "cb"}, []string{"ab", "c"}},
			ans1662{false},
		},

		{
			para1662{[]string{"abc", "d", "defg"}, []string{"abcddefg"}},
			ans1662{true},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1662, q.para1662
				(arrayStringsAreEqual(p.word1, p.word2))
			}
		}
	}
}

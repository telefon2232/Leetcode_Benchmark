package leetcode

import (
	"testing"
)

type question1143 struct {
	para1143
	ans1143
}

// para 是参数
// one 代表第一个参数
type para1143 struct {
	text1 string
	text2 string
}

// ans 是答案
// one 代表第一个答案
type ans1143 struct {
	one int
}

func Benchmark_Problem1143(b *testing.B) {

	qs := []question1143{

		{
			para1143{"abcde", "ace"},
			ans1143{3},
		},

		{
			para1143{"abc", "abc"},
			ans1143{3},
		},

		{
			para1143{"abc", "def"},
			ans1143{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1143, q.para1143
				(longestCommonSubsequence(p.text1, p.text2))
			}
		}
	}
}

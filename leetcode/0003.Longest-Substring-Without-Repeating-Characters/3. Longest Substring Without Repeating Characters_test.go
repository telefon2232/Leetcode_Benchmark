package leetcode

import (
	"testing"
)

type question3 struct {
	para3
	ans3
}

// para 是参数
// one 代表第一个参数
type para3 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans3 struct {
	one int
}

func Benchmark_Problem3(b *testing.B) {

	qs := []question3{

		{
			para3{"abcabcbb"},
			ans3{3},
		},

		{
			para3{"bbbbb"},
			ans3{1},
		},

		{
			para3{"pwwkew"},
			ans3{3},
		},

		{
			para3{""},
			ans3{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans3, q.para3
		(lengthOfLongestSubstring(p.s))
	}
}}}

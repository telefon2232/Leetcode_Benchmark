package leetcode

import (
	"testing"
)

type question1763 struct {
	para1763
	ans1763
}

// para 是参数
// one 代表第一个参数
type para1763 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans1763 struct {
	one string
}

func Benchmark_Problem1763(b *testing.B) {

	qs := []question1763{

		{
			para1763{"YazaAay"},
			ans1763{"aAa"},
		},

		{
			para1763{"Bb"},
			ans1763{"Bb"},
		},

		{
			para1763{"c"},
			ans1763{""},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans1763, q.para1763
		(longestNiceSubstring(p.s))
	}
}}}

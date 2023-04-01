package leetcode

import (
	"testing"
)

type question290 struct {
	para290
	ans290
}

// para 是参数
// one 代表第一个参数
type para290 struct {
	one string
	two string
}

// ans 是答案
// one 代表第一个答案
type ans290 struct {
	one bool
}

func Benchmark_Problem290(b *testing.B) {

	qs := []question290{

		{
			para290{"abba", "dog cat cat dog"},
			ans290{true},
		},

		{
			para290{"abba", "dog cat cat fish"},
			ans290{false},
		},

		{
			para290{"aaaa", "dog cat cat dog"},
			ans290{false},
		},

		{
			para290{"abba", "dog dog dog dog"},
			ans290{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans290, q.para290
		(wordPattern(p.one, p.two))
	}
}}}

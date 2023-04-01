package leetcode

import (
	"testing"
)

type question524 struct {
	para524
	ans524
}

// para 是参数
// one 代表第一个参数
type para524 struct {
	s   string
	one []string
}

// ans 是答案
// one 代表第一个答案
type ans524 struct {
	one string
}

func Benchmark_Problem524(b *testing.B) {

	qs := []question524{

		{
			para524{"abpcplea", []string{"ale", "apple", "monkey", "plea"}},
			ans524{"apple"},
		},

		{
			para524{"abpcplea", []string{"a", "b", "c"}},
			ans524{"a"},
		},

		{
			para524{"abpcplea", []string{"aaaaa", "b", "c"}},
			ans524{"b"},
		},

		{
			para524{"bab", []string{"ba", "ab", "a", "b"}},
			ans524{"ab"},
		},

		{
			para524{"aewfafwafjlwajflwajflwafj", []string{"apple", "ewaf", "awefawfwaf", "awef", "awefe", "ewafeffewafewf"}},
			ans524{"ewaf"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans524, q.para524
		(findLongestWord(p.s, p.one))
	}
}}}

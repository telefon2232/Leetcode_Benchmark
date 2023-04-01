package leetcode

import (
	"testing"
)

type question821 struct {
	para821
	ans821
}

// para 是参数
// one 代表第一个参数
type para821 struct {
	s string
	c byte
}

// ans 是答案
// one 代表第一个答案
type ans821 struct {
	one []int
}

func Benchmark_Problem821(b *testing.B) {

	qs := []question821{

		{
			para821{"loveleetcode", 'e'},
			ans821{[]int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}},
		},

		{
			para821{"baaa", 'b'},
			ans821{[]int{0, 1, 2, 3}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans821, q.para821
		(shortestToChar(p.s, p.c))
	}
}}}

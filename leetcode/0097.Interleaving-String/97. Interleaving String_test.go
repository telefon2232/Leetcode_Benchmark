package leetcode

import (
	"testing"
)

type question97 struct {
	para97
	ans97
}

// para 是参数
// one 代表第一个参数
type para97 struct {
	s1 string
	s2 string
	s3 string
}

// ans 是答案
// one 代表第一个答案
type ans97 struct {
	one bool
}

func Benchmark_Problem97(b *testing.B) {

	qs := []question97{

		{
			para97{"aabcc", "dbbca", "aadbbcbcac"},
			ans97{true},
		},

		{
			para97{"aabcc", "dbbca", "aadbbbaccc"},
			ans97{false},
		},

		{
			para97{"", "", ""},
			ans97{true},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans97, q.para97
		(isInterleave(p.s1, p.s2, p.s3))
	}
}}}

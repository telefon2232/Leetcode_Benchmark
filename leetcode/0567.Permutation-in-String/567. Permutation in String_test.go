package leetcode

import (
	"testing"
)

type question567 struct {
	para567
	ans567
}

// para 是参数
// one 代表第一个参数
type para567 struct {
	s string
	p string
}

// ans 是答案
// one 代表第一个答案
type ans567 struct {
	one bool
}

func Benchmark_Problem567(b *testing.B) {

	qs := []question567{

		{
			para567{"ab", "abab"},
			ans567{true},
		},

		{
			para567{"abc", "cbaebabacd"},
			ans567{true},
		},

		{
			para567{"abc", ""},
			ans567{false},
		},

		{
			para567{"abc", "abacbabc"},
			ans567{true},
		},

		{
			para567{"ab", "eidboaoo"},
			ans567{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans567, q.para567
		(checkInclusion(p.s, p.p))
	}
}}}

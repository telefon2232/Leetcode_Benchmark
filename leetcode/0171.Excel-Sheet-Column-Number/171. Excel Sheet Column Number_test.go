package leetcode

import (
	"testing"
)

type question171 struct {
	para171
	ans171
}

// para 是参数
// one 代表第一个参数
type para171 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans171 struct {
	one int
}

func Benchmark_Problem171(b *testing.B) {

	qs := []question171{

		{
			para171{"A"},
			ans171{1},
		},

		{
			para171{"AB"},
			ans171{28},
		},

		{
			para171{"ZY"},
			ans171{701},
		},

		{
			para171{"ABC"},
			ans171{731},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans171, q.para171
		(titleToNumber(p.s))
	}
}}}

package leetcode

import (
	"testing"
)

type question299 struct {
	para299
	ans299
}

// para 是参数
type para299 struct {
	secret string
	guess  string
}

// ans 是答案
type ans299 struct {
	ans string
}

func Benchmark_Problem299(b *testing.B) {

	qs := []question299{

		{
			para299{"1807", "7810"},
			ans299{"1A3B"},
		},

		{
			para299{"1123", "0111"},
			ans299{"1A1B"},
		},

		{
			para299{"1", "0"},
			ans299{"0A0B"},
		},

		{
			para299{"1", "1"},
			ans299{"1A0B"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans299, q.para299
		(getHint(p.secret, p.guess))
	}
}}}

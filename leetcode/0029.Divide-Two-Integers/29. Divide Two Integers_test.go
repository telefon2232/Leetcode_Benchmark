package leetcode

import (
	"testing"
)

type question29 struct {
	para29
	ans29
}

// para 是参数
// one 代表第一个参数
type para29 struct {
	dividend int
	divisor  int
}

// ans 是答案
// one 代表第一个答案
type ans29 struct {
	one int
}

func Benchmark_Problem29(b *testing.B) {

	qs := []question29{

		{
			para29{10, 3},
			ans29{3},
		},

		{
			para29{7, -3},
			ans29{-2},
		},

		{
			para29{-1, 1},
			ans29{-1},
		},

		{
			para29{1, -1},
			ans29{-1},
		},

		{
			para29{2147483647, 3},
			ans29{715827882},
		},
	}


	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans29, q.para29
		(divide(p.dividend, p.divisor))
	}
}}}

package leetcode

import (
	"testing"
)

type question991 struct {
	para991
	ans991
}

// para 是参数
// one 代表第一个参数
type para991 struct {
	X int
	Y int
}

// ans 是答案
// one 代表第一个答案
type ans991 struct {
	one int
}

func Benchmark_Problem991(b *testing.B) {

	qs := []question991{

		{
			para991{2, 3},
			ans991{2},
		},

		{
			para991{5, 8},
			ans991{2},
		},

		{
			para991{3, 10},
			ans991{3},
		},

		{
			para991{1024, 1},
			ans991{1023},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans991, q.para991
		(brokenCalc(p.X, p.Y))
	}
}}}

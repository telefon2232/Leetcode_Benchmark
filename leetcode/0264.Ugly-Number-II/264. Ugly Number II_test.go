package leetcode

import (
	"testing"
)

type question264 struct {
	para264
	ans264
}

// para 是参数
// one 代表第一个参数
type para264 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans264 struct {
	one int
}

func Benchmark_Problem264(b *testing.B) {

	qs := []question264{

		{
			para264{10},
			ans264{12},
		},

		{
			para264{1},
			ans264{1},
		},

		{
			para264{6},
			ans264{6},
		},

		{
			para264{8},
			ans264{9},
		},

		{
			para264{14},
			ans264{20},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans264, q.para264
		(nthUglyNumber(p.one))
	}
}}}

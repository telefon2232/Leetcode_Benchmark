package leetcode

import (
	"testing"
)

type question42 struct {
	para42
	ans42
}

// para 是参数
// one 代表第一个参数
type para42 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans42 struct {
	one int
}

func Benchmark_Problem42(b *testing.B) {

	qs := []question42{

		{
			para42{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
			ans42{6},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans42, q.para42
		(trap(p.one))
	}
}}}

package leetcode

import (
	"testing"
)

type question169 struct {
	para169
	ans169
}

// para 是参数
// one 代表第一个参数
type para169 struct {
	s []int
}

// ans 是答案
// one 代表第一个答案
type ans169 struct {
	one int
}

func Benchmark_Problem169(b *testing.B) {

	qs := []question169{

		{
			para169{[]int{2, 2, 1}},
			ans169{2},
		},

		{
			para169{[]int{3, 2, 3}},
			ans169{3},
		},

		{
			para169{[]int{2, 2, 1, 1, 1, 2, 2}},
			ans169{2},
		},

		{
			para169{[]int{-2147483648}},
			ans169{-2147483648},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans169, q.para169
		(majorityElement(p.s))
	}
}}}

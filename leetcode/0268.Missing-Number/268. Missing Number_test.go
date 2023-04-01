package leetcode

import (
	"testing"
)

type question268 struct {
	para268
	ans268
}

// para 是参数
// one 代表第一个参数
type para268 struct {
	s []int
}

// ans 是答案
// one 代表第一个答案
type ans268 struct {
	one int
}

func Benchmark_Problem268(b *testing.B) {

	qs := []question268{

		{
			para268{[]int{3, 0, 1}},
			ans268{2},
		},

		{
			para268{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}},
			ans268{8},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans268, q.para268
		(missingNumber(p.s))
	}
}}}

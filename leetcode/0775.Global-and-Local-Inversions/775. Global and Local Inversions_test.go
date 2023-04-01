package leetcode

import (
	"testing"
)

type question775 struct {
	para775
	ans775
}

// para 是参数
// one 代表第一个参数
type para775 struct {
	A []int
}

// ans 是答案
// one 代表第一个答案
type ans775 struct {
	one bool
}

func Benchmark_Problem775(b *testing.B) {

	qs := []question775{
		{
			para775{[]int{1, 0, 2}},
			ans775{true},
		},

		{
			para775{[]int{1, 2, 0}},
			ans775{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans775, q.para775
		(isIdealPermutation(p.A))
	}
}}}

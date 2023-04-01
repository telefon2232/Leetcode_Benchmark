package leetcode

import (
	"testing"
)

type question377 struct {
	para377
	ans377
}

// para 是参数
// one 代表第一个参数
type para377 struct {
	n []int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans377 struct {
	one int
}

func Benchmark_Problem377(b *testing.B) {

	qs := []question377{

		{
			para377{[]int{1, 2, 3}, 4},
			ans377{7},
		},

		{
			para377{[]int{1, 2, 3}, 32},
			ans377{181997601},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans377, q.para377
		(combinationSum4(p.n, p.k))
	}
}}}

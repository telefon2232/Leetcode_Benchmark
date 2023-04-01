package leetcode

import (
	"testing"
)

type question667 struct {
	para667
	ans667
}

// para 是参数
// one 代表第一个参数
type para667 struct {
	n int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans667 struct {
	one []int
}

func Benchmark_Problem667(b *testing.B) {

	qs := []question667{

		{
			para667{3, 1},
			ans667{[]int{1, 2, 3}},
		},

		{
			para667{3, 2},
			ans667{[]int{1, 3, 2}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans667, q.para667
		(constructArray(p.n, p.k))
	}
}}}

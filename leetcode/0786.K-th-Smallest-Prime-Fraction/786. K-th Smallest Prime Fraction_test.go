package leetcode

import (
	"testing"
)

type question786 struct {
	para786
	ans786
}

// para 是参数
// one 代表第一个参数
type para786 struct {
	A []int
	K int
}

// ans 是答案
// one 代表第一个答案
type ans786 struct {
	one []int
}

func Benchmark_Problem786(b *testing.B) {

	qs := []question786{

		{
			para786{[]int{1, 2, 3, 5}, 3},
			ans786{[]int{2, 5}},
		},

		{
			para786{[]int{1, 7}, 1},
			ans786{[]int{1, 7}},
		},

		{
			para786{[]int{1, 2}, 1},
			ans786{[]int{1, 2}},
		},

		{
			para786{[]int{1, 2, 3, 5, 7}, 6},
			ans786{[]int{3, 7}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans786, q.para786
		(kthSmallestPrimeFraction(p.A, p.K))
	}
}}}

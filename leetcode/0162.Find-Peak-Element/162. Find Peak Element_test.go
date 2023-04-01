package leetcode

import (
	"testing"
)

type question162 struct {
	para162
	ans162
}

// para 是参数
// one 代表第一个参数
type para162 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans162 struct {
	one int
}

func Benchmark_Problem162(b *testing.B) {

	qs := []question162{

		{
			para162{[]int{2, 1, 2}},
			ans162{0},
		},

		{
			para162{[]int{3, 2, 1}},
			ans162{0},
		},

		{
			para162{[]int{1, 2}},
			ans162{1},
		},

		{
			para162{[]int{2, 1}},
			ans162{0},
		},

		{
			para162{[]int{1}},
			ans162{0},
		},

		{
			para162{[]int{1, 2, 3, 1}},
			ans162{2},
		},

		{
			para162{[]int{1, 2, 1, 3, 5, 6, 4}},
			ans162{5},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans162, q.para162
		(findPeakElement(p.one))
	}
}}}

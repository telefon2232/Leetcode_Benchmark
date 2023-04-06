package leetcode

import (
	"testing"
)

type question324 struct {
	para324
	ans324
}

// para 是参数
// one 代表第一个参数
type para324 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans324 struct {
	one []int
}

func Benchmark_Problem324(b *testing.B) {

	qs := []question324{

		{
			para324{[]int{}},
			ans324{[]int{}},
		},

		{
			para324{[]int{1}},
			ans324{[]int{1}},
		},

		{
			para324{[]int{1, 5, 1, 1, 6, 4}},
			ans324{[]int{1, 4, 1, 5, 1, 6}},
		},

		{
			para324{[]int{1, 3, 2, 2, 3, 1}},
			ans324{[]int{2, 3, 1, 3, 1, 2}},
		},

		{
			para324{[]int{1, 1, 2, 1, 2, 2, 1}},
			ans324{[]int{1, 2, 1, 2, 1, 2, 1}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans324, q.para324

				wiggleSort(p.one)

			}
		}
	}
}

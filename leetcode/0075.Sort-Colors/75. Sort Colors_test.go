package leetcode

import (
	"testing"
)

type question75 struct {
	para75
	ans75
}

// para 是参数
// one 代表第一个参数
type para75 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans75 struct {
	one []int
}

func Benchmark_Problem75(b *testing.B) {

	qs := []question75{

		{
			para75{[]int{}},
			ans75{[]int{}},
		},

		{
			para75{[]int{1}},
			ans75{[]int{1}},
		},

		{
			para75{[]int{2, 0, 2, 1, 1, 0}},
			ans75{[]int{0, 0, 1, 1, 2, 2}},
		},

		{
			para75{[]int{2, 0, 1, 1, 2, 0, 2, 1, 2, 0, 0, 0, 1, 2, 2, 2, 0, 1, 1}},
			ans75{[]int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans75, q.para75

				sortColors(p.one)

			}
		}
	}
}

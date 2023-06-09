package leetcode

import (
	"testing"
)

type question976 struct {
	para976
	ans976
}

// para 是参数
// one 代表第一个参数
type para976 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans976 struct {
	one int
}

func Benchmark_Problem976(b *testing.B) {

	qs := []question976{

		{
			para976{[]int{1, 2}},
			ans976{0},
		},
		{
			para976{[]int{1, 2, 3}},
			ans976{0},
		},

		{
			para976{[]int{}},
			ans976{0},
		},

		{
			para976{[]int{2, 1, 2}},
			ans976{5},
		},

		{
			para976{[]int{1, 1, 2}},
			ans976{0},
		},

		{
			para976{[]int{3, 2, 3, 4}},
			ans976{10},
		},

		{
			para976{[]int{3, 6, 2, 3}},
			ans976{8},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans976, q.para976
				(largestPerimeter(p.one))
			}
		}
	}
}

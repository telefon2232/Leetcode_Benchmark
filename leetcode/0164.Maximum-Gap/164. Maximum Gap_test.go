package leetcode

import (
	"testing"
)

type question164 struct {
	para164
	ans164
}

// para 是参数
// one 代表第一个参数
type para164 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans164 struct {
	one int
}

func Benchmark_Problem164(b *testing.B) {

	qs := []question164{

		{
			para164{[]int{3, 6, 9, 1}},
			ans164{3},
		},
		{
			para164{[]int{1}},
			ans164{0},
		},

		{
			para164{[]int{}},
			ans164{0},
		},

		{
			para164{[]int{2, 10}},
			ans164{8},
		},

		{
			para164{[]int{2, 435, 214, 64321, 643, 7234, 7, 436523, 7856, 8}},
			ans164{372202},
		},

		{
			para164{[]int{1, 10000000}},
			ans164{9999999},
		},
	}

	for _, q := range qs {
		_, p := q.ans164, q.para164
		(maximumGap1(p.one))
	}
}

package leetcode

import (
	"testing"
)

type question914 struct {
	para914
	ans914
}

// para 是参数
// one 代表第一个参数
type para914 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans914 struct {
	one bool
}

func Benchmark_Problem914(b *testing.B) {

	qs := []question914{

		{
			para914{[]int{1, 2, 3, 4, 4, 3, 2, 1}},
			ans914{true},
		},

		{
			para914{[]int{1, 1, 1, 2, 2, 2, 3, 3}},
			ans914{false},
		},

		{
			para914{[]int{1}},
			ans914{false},
		},

		{
			para914{[]int{1, 1}},
			ans914{true},
		},

		{
			para914{[]int{1, 1, 2, 2, 2, 2}},
			ans914{true},
		},
	}

	for _, q := range qs {
		_, p := q.ans914, q.para914
		(hasGroupsSizeX(p.one))
	}
}

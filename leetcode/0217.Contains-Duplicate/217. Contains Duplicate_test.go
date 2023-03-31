package leetcode

import (
	"testing"
)

type question217 struct {
	para217
	ans217
}

// para 是参数
// one 代表第一个参数
type para217 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans217 struct {
	one bool
}

func Benchmark_Problem217(b *testing.B) {

	qs := []question217{

		{
			para217{[]int{1, 2, 3, 1}},
			ans217{true},
		},

		{
			para217{[]int{1, 2, 3, 4}},
			ans217{false},
		},

		{
			para217{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
			ans217{true},
		},
	}

	for _, q := range qs {
		_, p := q.ans217, q.para217
		(containsDuplicate(p.one))
	}
}

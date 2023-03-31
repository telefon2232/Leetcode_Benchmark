package leetcode

import (
	"testing"
)

type question136 struct {
	para136
	ans136
}

// para 是参数
// one 代表第一个参数
type para136 struct {
	s []int
}

// ans 是答案
// one 代表第一个答案
type ans136 struct {
	one int
}

func Benchmark_Problem136(b *testing.B) {

	qs := []question136{

		{
			para136{[]int{2, 2, 1}},
			ans136{1},
		},

		{
			para136{[]int{4, 1, 2, 1, 2}},
			ans136{4},
		},
	}

	for _, q := range qs {
		_, p := q.ans136, q.para136
		(singleNumber(p.s))
	}
}

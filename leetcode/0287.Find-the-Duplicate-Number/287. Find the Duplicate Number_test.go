package leetcode

import (
	"testing"
)

type question287 struct {
	para287
	ans287
}

// para 是参数
// one 代表第一个参数
type para287 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans287 struct {
	one int
}

func Benchmark_Problem287(b *testing.B) {

	qs := []question287{

		{
			para287{[]int{1, 3, 4, 2, 2}},
			ans287{2},
		},

		{
			para287{[]int{3, 1, 3, 4, 2}},
			ans287{3},
		},

		{
			para287{[]int{2, 2, 2, 2, 2}},
			ans287{2},
		},
	}

	for _, q := range qs {
		_, p := q.ans287, q.para287
		(findDuplicate(p.one))
	}
}

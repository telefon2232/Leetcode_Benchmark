package leetcode

import (
	"testing"
)

type question845 struct {
	para845
	ans845
}

// para 是参数
// one 代表第一个参数
type para845 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans845 struct {
	one int
}

func Benchmark_Problem845(b *testing.B) {

	qs := []question845{
		{
			para845{[]int{875, 884, 239, 731, 723, 685}},
			ans845{4},
		},

		{
			para845{[]int{0, 1, 2, 3, 4, 5, 4, 3, 2, 1, 0}},
			ans845{11},
		},

		{
			para845{[]int{2, 3}},
			ans845{0},
		},

		{
			para845{[]int{2, 1, 4, 7, 3, 2, 5}},
			ans845{5},
		},

		{
			para845{[]int{2, 2, 2}},
			ans845{0},
		},
	}

	for _, q := range qs {
		_, p := q.ans845, q.para845
		(longestMountain(p.one))
	}
}

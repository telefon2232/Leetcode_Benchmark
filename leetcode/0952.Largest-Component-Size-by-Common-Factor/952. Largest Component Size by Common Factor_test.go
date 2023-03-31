package leetcode

import (
	"testing"
)

type question952 struct {
	para952
	ans952
}

// para 是参数
// one 代表第一个参数
type para952 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans952 struct {
	one int
}

func Benchmark_Problem952(b *testing.B) {

	qs := []question952{
		{
			para952{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			ans952{6},
		},

		{
			para952{[]int{4, 6, 15, 35}},
			ans952{4},
		},

		{
			para952{[]int{20, 50, 9, 63}},
			ans952{2},
		},

		{
			para952{[]int{2, 3, 6, 7, 4, 12, 21, 39}},
			ans952{8},
		},
	}

	for _, q := range qs {
		_, p := q.ans952, q.para952
		(largestComponentSize(p.one))
	}
}

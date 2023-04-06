package leetcode

import (
	"testing"
)

type question413 struct {
	para413
	ans413
}

// para 是参数
// one 代表第一个参数
type para413 struct {
	A []int
}

// ans 是答案
// one 代表第一个答案
type ans413 struct {
	one int
}

func Benchmark_Problem413(b *testing.B) {

	qs := []question413{

		{
			para413{[]int{1, 2, 3, 4}},
			ans413{3},
		},

		{
			para413{[]int{1, 2, 3, 4, 9}},
			ans413{3},
		},

		{
			para413{[]int{1, 2, 3, 4, 5, 6, 7}},
			ans413{3},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans413, q.para413
				(numberOfArithmeticSlices(p.A))
			}
		}
	}
}

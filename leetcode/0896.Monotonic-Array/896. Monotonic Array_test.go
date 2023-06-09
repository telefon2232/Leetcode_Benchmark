package leetcode

import (
	"testing"
)

type question896 struct {
	para896
	ans896
}

// para 是参数
// one 代表第一个参数
type para896 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans896 struct {
	one bool
}

func Benchmark_Problem896(b *testing.B) {

	qs := []question896{

		{
			para896{[]int{2, 1, 3}},
			ans896{false},
		},

		{
			para896{[]int{1, 2, 2, 3}},
			ans896{true},
		},

		{
			para896{[]int{6, 5, 4, 4}},
			ans896{true},
		},

		{
			para896{[]int{1, 3, 2}},
			ans896{false},
		},

		{
			para896{[]int{1, 2, 4, 5}},
			ans896{true},
		},

		{
			para896{[]int{1, 1, 1}},
			ans896{true},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans896, q.para896
				(isMonotonic(p.one))
			}
		}
	}
}

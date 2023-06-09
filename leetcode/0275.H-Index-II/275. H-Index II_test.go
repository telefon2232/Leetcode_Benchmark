package leetcode

import (
	"testing"
)

type question275 struct {
	para275
	ans275
}

// para 是参数
// one 代表第一个参数
type para275 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans275 struct {
	one int
}

func Benchmark_Problem275(b *testing.B) {

	qs := []question275{

		{
			para275{[]int{3, 6, 9, 1}},
			ans275{3},
		},
		{
			para275{[]int{1}},
			ans275{1},
		},

		{
			para275{[]int{}},
			ans275{0},
		},

		{
			para275{[]int{3, 0, 6, 1, 5}},
			ans275{3},
		},

		{
			para275{[]int{0, 1, 3, 5, 6}},
			ans275{3},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans275, q.para275
				(hIndex275(p.one))
			}
		}
	}
}

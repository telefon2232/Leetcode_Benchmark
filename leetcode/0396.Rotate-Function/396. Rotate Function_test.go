package leetcode

import (
	"testing"
)

type question396 struct {
	para396
	ans396
}

// para 是参数
// one 代表第一个参数
type para396 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans396 struct {
	one int
}

func Benchmark_Problem396(b *testing.B) {

	qs := []question396{
		{
			para396{[]int{4, 3, 2, 6}},
			ans396{26},
		},

		{
			para396{[]int{100}},
			ans396{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans396, q.para396
				(maxRotateFunction(p.one))
			}
		}
	}
}

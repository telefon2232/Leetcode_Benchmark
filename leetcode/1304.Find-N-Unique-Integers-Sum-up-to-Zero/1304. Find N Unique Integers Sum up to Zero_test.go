package leetcode

import (
	"testing"
)

type question1304 struct {
	para1304
	ans1304
}

// para 是参数
// one 代表第一个参数
type para1304 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans1304 struct {
	one []int
}

func Benchmark_Problem1304(b *testing.B) {

	qs := []question1304{

		{
			para1304{5},
			ans1304{[]int{-7, -1, 1, 3, 4}},
		},

		{
			para1304{0},
			ans1304{[]int{}},
		},

		{
			para1304{3},
			ans1304{[]int{-1, 0, 1}},
		},

		{
			para1304{1},
			ans1304{[]int{0}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1304, q.para1304
				(sumZero(p.one))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question735 struct {
	para735
	ans735
}

// para 是参数
// one 代表第一个参数
type para735 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans735 struct {
	one []int
}

func Benchmark_Problem735(b *testing.B) {

	qs := []question735{

		{
			para735{[]int{-1, -1, 1, -2}},
			ans735{[]int{-1, -1, -2}},
		},

		{
			para735{[]int{5, 10, -5}},
			ans735{[]int{5, 10}},
		},

		{
			para735{[]int{5, 8, -8}},
			ans735{[]int{5}},
		},

		{
			para735{[]int{8, -8}},
			ans735{[]int{}},
		},

		{
			para735{[]int{10, 2, -5}},
			ans735{[]int{10}},
		},

		{
			para735{[]int{-2, -1, 1, 2}},
			ans735{[]int{-2, -1, 1, 2}},
		},
	}

	for _, q := range qs {
		_, p := q.ans735, q.para735
		(asteroidCollision(p.one))
	}
}

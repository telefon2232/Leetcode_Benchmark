package leetcode

import (
	"testing"
)

type question457 struct {
	para457
	ans457
}

// para 是参数
// one 代表第一个参数
type para457 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans457 struct {
	one bool
}

func Benchmark_Problem457(b *testing.B) {

	qs := []question457{

		{
			para457{[]int{-1}},
			ans457{false},
		},

		{
			para457{[]int{3, 1, 2}},
			ans457{true},
		},

		{
			para457{[]int{-8, -1, 1, 7, 2}},
			ans457{false},
		},

		{
			para457{[]int{-1, -2, -3, -4, -5}},
			ans457{false},
		},

		{
			para457{[]int{}},
			ans457{false},
		},

		{
			para457{[]int{2, -1, 1, 2, 2}},
			ans457{true},
		},

		{
			para457{[]int{-1, 2}},
			ans457{false},
		},
		{
			para457{[]int{-2, 1, -1, -2, -2}},
			ans457{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans457, q.para457
		(circularArrayLoop(p.one))
	}
}}}

package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question92 struct {
	para92
	ans92
}

// para 是参数
// one 代表第一个参数
type para92 struct {
	one  []int
	m, n int
}

// ans 是答案
// one 代表第一个答案
type ans92 struct {
	one []int
}

func Benchmark_Problem92(b *testing.B) {

	qs := []question92{

		{
			para92{[]int{1, 2, 3, 4, 5}, 2, 4},
			ans92{[]int{1, 4, 3, 2, 5}},
		},

		{
			para92{[]int{1, 2, 3, 4, 5}, 2, 2},
			ans92{[]int{1, 2, 3, 4, 5}},
		},

		{
			para92{[]int{1, 2, 3, 4, 5}, 1, 5},
			ans92{[]int{5, 4, 3, 2, 1}},
		},

		{
			para92{[]int{1, 2, 3, 4, 5, 6}, 3, 4},
			ans92{[]int{1, 2, 4, 3, 5, 6}},
		},

		{
			para92{[]int{3, 5}, 1, 2},
			ans92{[]int{5, 3}},
		},

		{
			para92{[]int{3}, 3, 5},
			ans92{[]int{3}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans92, q.para92
		(structures.List2Ints(reverseBetween(structures.Ints2List(p.one), p.m, p.n)))
	}
}}}

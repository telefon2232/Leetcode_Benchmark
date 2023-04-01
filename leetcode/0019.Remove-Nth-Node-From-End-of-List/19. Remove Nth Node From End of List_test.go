package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question19 struct {
	para19
	ans19
}

// para 是参数
// one 代表第一个参数
type para19 struct {
	one []int
	n   int
}

// ans 是答案
// one 代表第一个答案
type ans19 struct {
	one []int
}

func Benchmark_Problem19(b *testing.B) {

	qs := []question19{

		{
			para19{[]int{1}, 3},
			ans19{[]int{1}},
		},

		{
			para19{[]int{1, 2}, 2},
			ans19{[]int{2}},
		},

		{
			para19{[]int{1}, 1},
			ans19{[]int{}},
		},

		{
			para19{[]int{1, 2, 3, 4, 5}, 10},
			ans19{[]int{1, 2, 3, 4, 5}},
		},

		{
			para19{[]int{1, 2, 3, 4, 5}, 1},
			ans19{[]int{1, 2, 3, 4}},
		},

		{
			para19{[]int{1, 2, 3, 4, 5}, 2},
			ans19{[]int{1, 2, 3, 5}},
		},

		{
			para19{[]int{1, 2, 3, 4, 5}, 3},
			ans19{[]int{1, 2, 4, 5}},
		},
		{
			para19{[]int{1, 2, 3, 4, 5}, 4},
			ans19{[]int{1, 3, 4, 5}},
		},

		{
			para19{[]int{1, 2, 3, 4, 5}, 5},
			ans19{[]int{2, 3, 4, 5}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans19, q.para19
		(structures.List2Ints(removeNthFromEnd(structures.Ints2List(p.one), p.n)))
	}
}}}

package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question86 struct {
	para86
	ans86
}

// para 是参数
// one 代表第一个参数
type para86 struct {
	one []int
	x   int
}

// ans 是答案
// one 代表第一个答案
type ans86 struct {
	one []int
}

func Benchmark_Problem86(b *testing.B) {

	qs := []question86{

		{
			para86{[]int{1, 4, 3, 2, 5, 2}, 3},
			ans86{[]int{1, 2, 2, 4, 3, 5}},
		},

		{
			para86{[]int{1, 1, 2, 2, 3, 3, 3}, 2},
			ans86{[]int{1, 1, 2, 2, 3, 3, 3}},
		},

		{
			para86{[]int{1, 4, 3, 2, 5, 2}, 0},
			ans86{[]int{1, 4, 3, 2, 5, 2}},
		},

		{
			para86{[]int{4, 3, 2, 5, 2}, 3},
			ans86{[]int{2, 2, 4, 3, 5}},
		},

		{
			para86{[]int{1, 1, 1, 1, 1, 1}, 1},
			ans86{[]int{1, 1, 1, 1, 1, 1}},
		},

		{
			para86{[]int{3, 1}, 2},
			ans86{[]int{1, 3}},
		},

		{
			para86{[]int{1, 2}, 3},
			ans86{[]int{1, 2}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans86, q.para86
				(structures.List2Ints(partition(structures.Ints2List(p.one), p.x)))
			}
		}
	}
}

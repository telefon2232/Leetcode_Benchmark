package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question143 struct {
	para143
	ans143
}

// para 是参数
// one 代表第一个参数
type para143 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans143 struct {
	one []int
}

func Benchmark_Problem143(b *testing.B) {

	qs := []question143{

		{
			para143{[]int{1, 2, 3, 4, 5}},
			ans143{[]int{1, 5, 2, 4, 3}},
		},
		{
			para143{[]int{1, 2, 3, 4}},
			ans143{[]int{1, 4, 2, 3}},
		},

		{
			para143{[]int{1}},
			ans143{[]int{1}},
		},

		{
			para143{[]int{}},
			ans143{[]int{}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans143, q.para143
				(structures.List2Ints(reorderList(structures.Ints2List(p.one))))
			}
		}
	}
}

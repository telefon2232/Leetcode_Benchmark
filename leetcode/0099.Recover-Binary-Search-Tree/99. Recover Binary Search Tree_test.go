package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question99 struct {
	para99
	ans99
}

// para 是参数
// one 代表第一个参数
type para99 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans99 struct {
	one []int
}

func Benchmark_Problem99(b *testing.B) {

	qs := []question99{

		{
			para99{[]int{1, 3, structures.NULL, structures.NULL, 2}},
			ans99{[]int{3, 1, structures.NULL, structures.NULL, 2}},
		},

		{
			para99{[]int{3, 1, 4, structures.NULL, structures.NULL, 2}},
			ans99{[]int{2, 1, 4, structures.NULL, structures.NULL, 3}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans99, q.para99

				rootOne := structures.Ints2TreeNode(p.one)
				recoverTree(rootOne)

			}
		}
	}
}

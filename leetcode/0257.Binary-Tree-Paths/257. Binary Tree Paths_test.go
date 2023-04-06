package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question257 struct {
	para257
	ans257
}

// para 是参数
// one 代表第一个参数
type para257 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans257 struct {
	one []string
}

func Benchmark_Problem257(b *testing.B) {

	qs := []question257{

		{
			para257{[]int{}},
			ans257{[]string{""}},
		},

		{
			para257{[]int{1, 2, 3, structures.NULL, 5, structures.NULL, structures.NULL}},
			ans257{[]string{"1->2->5", "1->3"}},
		},

		{
			para257{[]int{1, 2, 3, 4, 5, 6, 7}},
			ans257{[]string{"1->2->4", "1->2->5", "1->2->4", "1->2->5", "1->3->6", "1->3->7"}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans257, q.para257

				root := structures.Ints2TreeNode(p.one)
				(binaryTreePaths(root))
			}
		}
	}
}

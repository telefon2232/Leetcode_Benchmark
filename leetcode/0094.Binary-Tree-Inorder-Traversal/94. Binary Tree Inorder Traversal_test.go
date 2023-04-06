package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question94 struct {
	para94
	ans94
}

// para 是参数
// one 代表第一个参数
type para94 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans94 struct {
	one []int
}

func Benchmark_Problem94(b *testing.B) {

	qs := []question94{

		{
			para94{[]int{}},
			ans94{[]int{}},
		},

		{
			para94{[]int{1}},
			ans94{[]int{1}},
		},

		{
			para94{[]int{1, structures.NULL, 2, 3}},
			ans94{[]int{1, 2, 3}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans94, q.para94

				root := structures.Ints2TreeNode(p.one)
				(inorderTraversal(root))
			}
		}
	}
}

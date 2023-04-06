package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question235 struct {
	para235
	ans235
}

// para 是参数
// one 代表第一个参数
type para235 struct {
	one []int
	two []int
	thr []int
}

// ans 是答案
// one 代表第一个答案
type ans235 struct {
	one []int
}

func Benchmark_Problem235(b *testing.B) {

	qs := []question235{

		{
			para235{[]int{}, []int{}, []int{}},
			ans235{[]int{}},
		},

		{
			para235{[]int{6, 2, 8, 0, 4, 7, 9, structures.NULL, structures.NULL, 3, 5}, []int{2}, []int{8}},
			ans235{[]int{6}},
		},

		{
			para235{[]int{6, 2, 8, 0, 4, 7, 9, structures.NULL, structures.NULL, 3, 5}, []int{2}, []int{4}},
			ans235{[]int{2}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans235, q.para235

				rootOne := structures.Ints2TreeNode(p.one)
				rootTwo := structures.Ints2TreeNode(p.two)
				rootThr := structures.Ints2TreeNode(p.thr)
				(lowestCommonAncestor(rootOne, rootTwo, rootThr))
			}
		}
	}
}

package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question105 struct {
	para105
	ans105
}

// para 是参数
// one 代表第一个参数
type para105 struct {
	preorder []int
	inorder  []int
}

// ans 是答案
// one 代表第一个答案
type ans105 struct {
	one []int
}

func Benchmark_Problem105(b *testing.B) {

	qs := []question105{

		{
			para105{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}},
			ans105{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans105, q.para105

				(structures.Tree2ints(buildTree(p.preorder, p.inorder)))
			}
		}
	}
}

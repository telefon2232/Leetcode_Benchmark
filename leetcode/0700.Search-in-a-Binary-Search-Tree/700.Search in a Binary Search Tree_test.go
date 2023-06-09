package leetcode

import (
	"testing"
)

type question700 struct {
	para700
	ans700
}

// para 是参数
type para700 struct {
	root *TreeNode
	val  int
}

// ans 是答案
type ans700 struct {
	ans *TreeNode
}

func Benchmark_Problem700(b *testing.B) {

	qs := []question700{
		{
			para700{&TreeNode{Val: 4,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
				Right: &TreeNode{Val: 7, Left: nil, Right: nil}},
				2},
			ans700{&TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}},
		},

		{
			para700{&TreeNode{Val: 4,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
				Right: &TreeNode{Val: 7, Left: nil, Right: nil}},
				5},
			ans700{nil},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans700, q.para700
				(searchBST(p.root, p.val))
			}
		}
	}
}

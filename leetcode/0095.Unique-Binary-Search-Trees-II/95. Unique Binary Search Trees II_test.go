package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question95 struct {
	para95
	ans95
}

// para 是参数
// one 代表第一个参数
type para95 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans95 struct {
	one []*TreeNode
}

func Benchmark_Problem95(b *testing.B) {

	qs := []question95{

		{
			para95{1},
			ans95{[]*TreeNode{{Val: 1, Left: nil, Right: nil}}},
		},

		{
			para95{3},
			ans95{[]*TreeNode{
				{Val: 1, Left: nil, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: nil}},
				{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: nil, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}},
				{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: nil}, Right: nil},
				{Val: 3, Left: &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: nil, Right: nil}}, Right: nil},
				{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
			}}},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans95, q.para95

				trees := generateTrees(p.one)
				for _, t := range trees {
					(structures.Tree2Preorder(t))
				}
			}
		}
	}
}

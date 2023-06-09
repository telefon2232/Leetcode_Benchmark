package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question653 struct {
	para653
	ans653
}

// para 是参数
// one 代表第一个参数
type para653 struct {
	one []int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans653 struct {
	one bool
}

func Benchmark_Problem653(b *testing.B) {

	qs := []question653{

		{
			para653{[]int{}, 0},
			ans653{false},
		},

		{
			para653{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}, 29},
			ans653{true},
		},

		{
			para653{[]int{1, 2, 3, 4, structures.NULL, structures.NULL, 5}, 9},
			ans653{true},
		},

		{
			para653{[]int{1, 2, 3, 4, structures.NULL, 5}, 4},
			ans653{true},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans653, q.para653

				root := structures.Ints2TreeNode(p.one)
				(findTarget(root, p.k))
			}
		}
	}
}

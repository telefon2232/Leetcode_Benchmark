package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question437 struct {
	para437
	ans437
}

// para 是参数
// one 代表第一个参数
type para437 struct {
	one []int
	sum int
}

// ans 是答案
// one 代表第一个答案
type ans437 struct {
	one int
}

func Benchmark_Problem437(b *testing.B) {

	qs := []question437{
		{
			para437{[]int{}, 0},
			ans437{0},
		},

		{
			para437{[]int{10, 5, -3, 3, 2, structures.NULL, 11, 3, -2, structures.NULL, 1}, 8},
			ans437{3},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans437, q.para437

				root := structures.Ints2TreeNode(p.one)
				(pathSum(root, p.sum))
			}
		}
	}
}

package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question113 struct {
	para113
	ans113
}

// para 是参数
// one 代表第一个参数
type para113 struct {
	one []int
	sum int
}

// ans 是答案
// one 代表第一个答案
type ans113 struct {
	one [][]int
}

func Benchmark_Problem113(b *testing.B) {

	qs := []question113{
		{
			para113{[]int{1, 0, 1, 1, 2, 0, -1, 0, 1, -1, 0, -1, 0, 1, 0}, 2},
			ans113{[][]int{{1, 0, 1, 0}, {1, 0, 2, -1}, {1, 1, 0, 0}, {1, 1, -1, 1}}},
		},

		{
			para113{[]int{}, 0},
			ans113{[][]int{{}}},
		},

		{
			para113{[]int{5, 4, 8, 11, structures.NULL, 13, 4, 7, 2, structures.NULL, structures.NULL, 5, 1}, 22},
			ans113{[][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans113, q.para113

				root := structures.Ints2TreeNode(p.one)
				(pathSum(root, p.sum))
			}
		}
	}
}

package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question987 struct {
	para987
	ans987
}

// para 是参数
// one 代表第一个参数
type para987 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans987 struct {
	one [][]int
}

func Benchmark_Problem987(b *testing.B) {

	qs := []question987{

		{
			para987{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			ans987{[][]int{{9}, {3, 15}, {20}, {7}}},
		},

		{
			para987{[]int{1, 2, 3, 4, 5, 6, 7}},
			ans987{[][]int{{4}, {2}, {1, 5, 6}, {3}, {7}}},
		},

		{
			para987{[]int{1, 2, 3, 4, 6, 5, 7}},
			ans987{[][]int{{4}, {2}, {1, 5, 6}, {3}, {7}}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans987, q.para987

				root := structures.Ints2TreeNode(p.one)
				(verticalTraversal(root))
			}
		}
	}
}

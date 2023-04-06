package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question1110 struct {
	para1110
	ans1110
}

// para 是参数
// one 代表第一个参数
type para1110 struct {
	one []int
	two []int
}

// ans 是答案
// one 代表第一个答案
type ans1110 struct {
	one [][]int
}

func Benchmark_Problem1110(b *testing.B) {

	qs := []question1110{

		{
			para1110{[]int{1, 2, 3, 4, 5, 6, 7}, []int{3, 5}},
			ans1110{[][]int{{1, 2, structures.NULL, 4}, {6}, {7}}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1110, q.para1110
				tree := structures.Ints2TreeNode(p.one)
				(delNodes(tree, p.two))
			}
		}
	}
}

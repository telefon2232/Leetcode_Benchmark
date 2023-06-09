package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question199 struct {
	para199
	ans199
}

// para 是参数
// one 代表第一个参数
type para199 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans199 struct {
	one []int
}

func Benchmark_Problem199(b *testing.B) {

	qs := []question199{

		{
			para199{[]int{}},
			ans199{[]int{}},
		},

		{
			para199{[]int{1}},
			ans199{[]int{1}},
		},

		{
			para199{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			ans199{[]int{3, 20, 7}},
		},

		{
			para199{[]int{1, 2, 3, 4, structures.NULL, structures.NULL, 5}},
			ans199{[]int{1, 3, 5}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans199, q.para199

				root := structures.Ints2TreeNode(p.one)
				(rightSideView(root))
			}
		}
	}
}

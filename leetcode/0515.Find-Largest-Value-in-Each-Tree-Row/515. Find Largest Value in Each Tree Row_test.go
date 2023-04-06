package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question515 struct {
	para515
	ans515
}

// para 是参数
// one 代表第一个参数
type para515 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans515 struct {
	one []int
}

func Benchmark_Problem515(b *testing.B) {

	qs := []question515{

		{
			para515{[]int{}},
			ans515{[]int{}},
		},

		{
			para515{[]int{1}},
			ans515{[]int{1}},
		},

		{
			para515{[]int{1, 3, 2, 5, 3, structures.NULL, 9}},
			ans515{[]int{1, 3, 9}},
		},

		{
			para515{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			ans515{[]int{3, 20, 15}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans515, q.para515

				root := structures.Ints2TreeNode(p.one)
				(largestValues(root))
			}
		}
	}
}

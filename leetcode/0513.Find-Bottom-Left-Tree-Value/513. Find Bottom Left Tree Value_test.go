package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question513 struct {
	para513
	ans513
}

// para 是参数
// one 代表第一个参数
type para513 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans513 struct {
	one int
}

func Benchmark_Problem513(b *testing.B) {

	qs := []question513{

		{
			para513{[]int{}},
			ans513{0},
		},

		{
			para513{[]int{1}},
			ans513{1},
		},

		{
			para513{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			ans513{15},
		},

		{
			para513{[]int{1, 2, 3, 4, structures.NULL, structures.NULL, 5}},
			ans513{4},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans513, q.para513

				root := structures.Ints2TreeNode(p.one)
				(findBottomLeftValue(root))
			}
		}
	}
}

package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question572 struct {
	para572
	ans572
}

// para 是参数
// one 代表第一个参数
type para572 struct {
	s []int
	t []int
}

// ans 是答案
// one 代表第一个答案
type ans572 struct {
	one bool
}

func Benchmark_Problem572(b *testing.B) {

	qs := []question572{

		{
			para572{[]int{}, []int{}},
			ans572{true},
		},

		{
			para572{[]int{3, 4, 5, 1, 2}, []int{4, 1, 2}},
			ans572{true},
		},

		{
			para572{[]int{1, 1}, []int{1}},
			ans572{true},
		},

		{
			para572{[]int{1, structures.NULL, 1, structures.NULL, 1, structures.NULL, 1, structures.NULL, 1, structures.NULL, 1, structures.NULL, 1, structures.NULL, 1, structures.NULL, 1, structures.NULL, 1, structures.NULL, 1, 2}, []int{1, structures.NULL, 1, structures.NULL, 1, structures.NULL, 1, structures.NULL, 1, structures.NULL, 1, 2}},
			ans572{true},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans572, q.para572

				roots := structures.Ints2TreeNode(p.s)
				roott := structures.Ints2TreeNode(p.t)
				(isSubtree(roots, roott))
			}
		}
	}
}

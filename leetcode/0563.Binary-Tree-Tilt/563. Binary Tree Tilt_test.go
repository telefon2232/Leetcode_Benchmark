package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question563 struct {
	para563
	ans563
}

// para 是参数
// one 代表第一个参数
type para563 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans563 struct {
	one int
}

func Benchmark_Problem563(b *testing.B) {

	qs := []question563{

		{
			para563{[]int{}},
			ans563{0},
		},

		{
			para563{[]int{1}},
			ans563{0},
		},

		{
			para563{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			ans563{41},
		},

		{
			para563{[]int{1, 2, 3, 4, structures.NULL, structures.NULL, 5}},
			ans563{11},
		},

		{
			para563{[]int{1, 2, 3, 4, structures.NULL, 5}},
			ans563{11},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans563, q.para563

				root := structures.Ints2TreeNode(p.one)
				(findTilt(root))
			}
		}
	}
}

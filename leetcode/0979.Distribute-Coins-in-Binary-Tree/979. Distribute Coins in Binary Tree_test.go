package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question979 struct {
	para979
	ans979
}

// para 是参数
// one 代表第一个参数
type para979 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans979 struct {
	one int
}

func Benchmark_Problem979(b *testing.B) {

	qs := []question979{

		{
			para979{[]int{}},
			ans979{0},
		},

		{
			para979{[]int{1}},
			ans979{0},
		},

		{
			para979{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			ans979{41},
		},

		{
			para979{[]int{1, 2, 3, 4, structures.NULL, structures.NULL, 5}},
			ans979{11},
		},

		{
			para979{[]int{1, 2, 3, 4, structures.NULL, 5}},
			ans979{11},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans979, q.para979

				root := structures.Ints2TreeNode(p.one)
				(distributeCoins(root))
			}
		}
	}
}

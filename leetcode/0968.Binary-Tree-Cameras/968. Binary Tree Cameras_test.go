package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question968 struct {
	para968
	ans968
}

// para 是参数
// one 代表第一个参数
type para968 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans968 struct {
	one int
}

func Benchmark_Problem968(b *testing.B) {

	qs := []question968{

		{
			para968{[]int{0, 0, structures.NULL, 0, 0}},
			ans968{1},
		},

		{
			para968{[]int{0, 0, structures.NULL, 0, structures.NULL, 0, structures.NULL, structures.NULL, 0}},
			ans968{2},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans968, q.para968

				root := structures.Ints2TreeNode(p.one)
				(minCameraCover(root))
			}
		}
	}
}

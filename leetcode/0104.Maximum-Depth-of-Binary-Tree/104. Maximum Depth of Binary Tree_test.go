package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question104 struct {
	para104
	ans104
}

// para 是参数
// one 代表第一个参数
type para104 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans104 struct {
	one int
}

func Benchmark_Problem104(b *testing.B) {

	qs := []question104{

		{
			para104{[]int{}},
			ans104{0},
		},

		{
			para104{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			ans104{3},
		},
	}


	for _, q := range qs {
		_, p := q.ans104, q.para104
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", maxDepth(root))
	}
}

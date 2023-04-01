package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question145 struct {
	para145
	ans145
}

// para 是参数
// one 代表第一个参数
type para145 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans145 struct {
	one []int
}

func Benchmark_Problem145(b *testing.B) {

	qs := []question145{

		{
			para145{[]int{}},
			ans145{[]int{}},
		},

		{
			para145{[]int{1}},
			ans145{[]int{1}},
		},

		{
			para145{[]int{1, structures.NULL, 2, 3}},
			ans145{[]int{1, 2, 3}},
		},
	}


	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans145, q.para145
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", postorderTraversal(root))
	}
}}}

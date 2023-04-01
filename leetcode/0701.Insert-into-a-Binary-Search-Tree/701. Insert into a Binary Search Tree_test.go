package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question701 struct {
	para701
	ans701
}

// para 是参数
// one 代表第一个参数
type para701 struct {
	root []int
	val  int
}

// ans 是答案
// one 代表第一个答案
type ans701 struct {
	one []int
}

func Benchmark_Problem701(b *testing.B) {

	qs := []question701{

		{
			para701{[]int{4, 2, 7, 1, 3}, 5},
			ans701{[]int{4, 2, 7, 1, 3, 5}},
		},

		{
			para701{[]int{40, 20, 60, 10, 30, 50, 70}, 25},
			ans701{[]int{40, 20, 60, 10, 30, 50, 70, structures.NULL, structures.NULL, 25}},
		},

		{
			para701{[]int{4, 2, 7, 1, 3, structures.NULL, structures.NULL, structures.NULL, structures.NULL, structures.NULL, structures.NULL}, 5},
			ans701{[]int{4, 2, 7, 1, 3, 5}},
		},
	}


	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans701, q.para701
		fmt.Printf("【input】:%v      ", p)
		rootOne := structures.Ints2TreeNode(p.root)
		fmt.Printf("【output】:%v      \n", structures.Tree2ints(insertIntoBST(rootOne, p.val)))
	}
}}}

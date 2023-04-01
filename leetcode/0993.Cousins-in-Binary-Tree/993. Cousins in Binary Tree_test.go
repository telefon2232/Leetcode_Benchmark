package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question993 struct {
	para993
	ans993
}

// para 是参数
// one 代表第一个参数
type para993 struct {
	one []int
	x   int
	y   int
}

// ans 是答案
// one 代表第一个答案
type ans993 struct {
	one bool
}

func Benchmark_Problem993(b *testing.B) {

	qs := []question993{

		{
			para993{[]int{1, 2, 3, 4}, 4, 3},
			ans993{false},
		},

		{
			para993{[]int{1, 2, 3, structures.NULL, 4, structures.NULL, 5}, 5, 4},
			ans993{true},
		},

		{
			para993{[]int{1, 2, 3, structures.NULL, 4}, 2, 3},
			ans993{false},
		},
	}


	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans993, q.para993
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", isCousins(root, p.x, p.y))
	}
}}}

package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question971 struct {
	para971
	ans971
}

// para 是参数
// one 代表第一个参数
type para971 struct {
	one    []int
	voyage []int
}

// ans 是答案
// one 代表第一个答案
type ans971 struct {
	one []int
}

func Benchmark_Problem971(b *testing.B) {

	qs := []question971{

		{
			para971{[]int{1, 2, structures.NULL}, []int{2, 1}},
			ans971{[]int{-1}},
		},

		{
			para971{[]int{1, 2, 3}, []int{1, 3, 2}},
			ans971{[]int{1}},
		},

		{
			para971{[]int{1, 2, 3}, []int{1, 2, 3}},
			ans971{[]int{}},
		},
	}


	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans971, q.para971
		fmt.Printf("【input】:%v      ", p)
		rootOne := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", flipMatchVoyage(rootOne, p.voyage))
	}
}}}

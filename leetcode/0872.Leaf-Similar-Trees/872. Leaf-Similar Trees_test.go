package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question872 struct {
	para872
	ans872
}

// para 是参数
// one 代表第一个参数
type para872 struct {
	one []int
	two []int
}

// ans 是答案
// one 代表第一个答案
type ans872 struct {
	one bool
}

func Benchmark_Problem872(b *testing.B) {

	qs := []question872{

		{
			para872{[]int{-10, -3, 0, 5, 9}, []int{-10, -3, 0, 5, 9}},
			ans872{true},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans872, q.para872
				tree1 := structures.Ints2TreeNode(p.one)
				tree2 := structures.Ints2TreeNode(p.two)
				(leafSimilar(tree1, tree2))
			}
		}
	}
}

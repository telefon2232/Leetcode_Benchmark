package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question114 struct {
	para114
	ans114
}

// para 是参数
// one 代表第一个参数
type para114 struct {
	one []string
}

// ans 是答案
// one 代表第一个答案
type ans114 struct {
	one []string
}

func Benchmark_Problem114(b *testing.B) {

	qs := []question114{

		{
			para114{[]string{"1", "2", "5", "3", "4", "null", "6"}},
			ans114{[]string{"1", "null", "2", "null", "3", "null", "4", "null", "5", "null", "6"}},
		},

		{
			para114{[]string{"0"}},
			ans114{[]string{"0"}},
		},

		{
			para114{[]string{"1", "2", "3", "4", "5", "6"}},
			ans114{[]string{"1", "2", "4", "5", "3", "6", "null"}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans114, q.para114

				rootOne := structures.Strings2TreeNode(p.one)
				flatten(rootOne)
				(structures.Tree2LevelOrderStrings(rootOne))
				(structures.Tree2PreOrderStrings(rootOne))
			}
		}
	}
}

package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question101 struct {
	para101
	ans101
}

// para 是参数
// one 代表第一个参数
type para101 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans101 struct {
	one bool
}

func Benchmark_Problem101(b *testing.B) {

	qs := []question101{

		{
			para101{[]int{3, 4, 4, 5, structures.NULL, structures.NULL, 5, 6, structures.NULL, structures.NULL, 6}},
			ans101{true},
		},

		{
			para101{[]int{1, 2, 2, structures.NULL, 3, 3}},
			ans101{true},
		},

		{
			para101{[]int{}},
			ans101{true},
		},

		{
			para101{[]int{1}},
			ans101{true},
		},

		{
			para101{[]int{1, 2, 3}},
			ans101{false},
		},

		{
			para101{[]int{1, 2, 2, 3, 4, 4, 3}},
			ans101{true},
		},

		{
			para101{[]int{1, 2, 2, structures.NULL, 3, structures.NULL, 3}},
			ans101{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans101, q.para101

				rootOne := structures.Ints2TreeNode(p.one)
				(isSymmetric(rootOne))
			}
		}
	}
}

package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question876 struct {
	para876
	ans876
}

// para 是参数
// one 代表第一个参数
type para876 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans876 struct {
	one int
}

func Benchmark_Problem876(b *testing.B) {

	qs := []question876{

		{
			para876{[]int{1, 2, 3, 4, 5}},
			ans876{3},
		},
		{
			para876{[]int{1, 2, 3, 4}},
			ans876{3},
		},

		{
			para876{[]int{1}},
			ans876{1},
		},

		{
			para876{[]int{}},
			ans876{},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans876, q.para876
				(structures.List2Ints(middleNode(structures.Ints2List(p.one))))
			}
		}
	}
}

package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question1019 struct {
	para1019
	ans1019
}

// para 是参数
// one 代表第一个参数
type para1019 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans1019 struct {
	one []int
}

func Benchmark_Problem1019(b *testing.B) {

	qs := []question1019{

		{
			para1019{[]int{2, 1, 5}},
			ans1019{[]int{5, 5, 0}},
		},

		{
			para1019{[]int{2, 7, 4, 3, 5}},
			ans1019{[]int{7, 0, 5, 5, 0}},
		},

		{
			para1019{[]int{1, 7, 5, 1, 9, 2, 5, 1}},
			ans1019{[]int{7, 9, 9, 9, 0, 5, 0, 0}},
		},

		{
			para1019{[]int{1, 7, 5, 1, 9, 2, 5, 6, 7, 8, 1}},
			ans1019{[]int{7, 9, 9, 9, 0, 5, 6, 7, 8, 0, 0}},
		},
	}

	for _, q := range qs {
		_, p := q.ans1019, q.para1019
		(nextLargerNodes(structures.Ints2List(p.one)))
	}
}

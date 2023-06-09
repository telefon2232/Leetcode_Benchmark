package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question24 struct {
	para24
	ans24
}

// para 是参数
// one 代表第一个参数
type para24 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans24 struct {
	one []int
}

func Benchmark_Problem24(b *testing.B) {

	qs := []question24{

		{
			para24{[]int{}},
			ans24{[]int{}},
		},

		{
			para24{[]int{1}},
			ans24{[]int{1}},
		},

		{
			para24{[]int{1, 2, 3, 4}},
			ans24{[]int{2, 1, 4, 3}},
		},

		{
			para24{[]int{1, 2, 3, 4, 5}},
			ans24{[]int{2, 1, 4, 3, 5}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans24, q.para24
				(structures.List2Ints(swapPairs(structures.Ints2List(p.one))))
			}
		}
	}
}

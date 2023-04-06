package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question445 struct {
	para445
	ans445
}

// para 是参数
// one 代表第一个参数
type para445 struct {
	one     []int
	another []int
}

// ans 是答案
// one 代表第一个答案
type ans445 struct {
	one []int
}

func Benchmark_Problem445(b *testing.B) {

	qs := []question445{

		{
			para445{[]int{}, []int{}},
			ans445{[]int{}},
		},

		{
			para445{[]int{1}, []int{1}},
			ans445{[]int{2}},
		},

		{
			para445{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			ans445{[]int{2, 4, 6, 8}},
		},

		{
			para445{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
			ans445{[]int{2, 4, 6, 9, 0}},
		},

		{
			para445{[]int{1}, []int{9, 9, 9, 9, 9}},
			ans445{[]int{1, 0, 0, 0, 0, 0}},
		},

		{
			para445{[]int{9, 9, 9, 9, 9}, []int{1}},
			ans445{[]int{1, 0, 0, 0, 0, 0}},
		},

		{
			para445{[]int{2, 4, 3}, []int{5, 6, 4}},
			ans445{[]int{8, 0, 7}},
		},

		{
			para445{[]int{1, 8, 3}, []int{7, 1}},
			ans445{[]int{2, 5, 4}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans445, q.para445
				(structures.List2Ints(addTwoNumbers445(structures.Ints2List(p.one), structures.Ints2List(p.another))))
			}
		}
	}
}

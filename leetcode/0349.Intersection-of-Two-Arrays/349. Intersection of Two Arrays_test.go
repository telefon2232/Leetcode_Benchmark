package leetcode

import (
	"testing"
)

type question349 struct {
	para349
	ans349
}

// para 是参数
// one 代表第一个参数
type para349 struct {
	one     []int
	another []int
}

// ans 是答案
// one 代表第一个答案
type ans349 struct {
	one []int
}

func Benchmark_Problem349(b *testing.B) {

	qs := []question349{

		{
			para349{[]int{}, []int{}},
			ans349{[]int{}},
		},

		{
			para349{[]int{1}, []int{1}},
			ans349{[]int{1}},
		},

		{
			para349{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			ans349{[]int{1, 2, 3, 4}},
		},

		{
			para349{[]int{1, 2, 2, 1}, []int{2, 2}},
			ans349{[]int{2}},
		},

		{
			para349{[]int{1}, []int{9, 9, 9, 9, 9}},
			ans349{[]int{}},
		},

		{
			para349{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}},
			ans349{[]int{9, 4}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans349, q.para349
				(intersection(p.one, p.another))
			}
		}
	}
}

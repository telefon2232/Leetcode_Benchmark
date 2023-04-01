package leetcode

import (
	"testing"
)

type question496 struct {
	para496
	ans496
}

// para 是参数
// one 代表第一个参数
type para496 struct {
	one     []int
	another []int
}

// ans 是答案
// one 代表第一个答案
type ans496 struct {
	one []int
}

func Benchmark_Problem496(b *testing.B) {

	qs := []question496{

		{
			para496{[]int{4, 1, 2}, []int{1, 3, 4, 2}},
			ans496{[]int{-1, 3, -1}},
		},

		{
			para496{[]int{2, 4}, []int{1, 2, 3, 4}},
			ans496{[]int{3, -1}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans496, q.para496
		(nextGreaterElement(p.one, p.another))
	}
}}}

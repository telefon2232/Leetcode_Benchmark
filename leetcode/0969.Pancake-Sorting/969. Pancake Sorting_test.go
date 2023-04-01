package leetcode

import (
	"testing"
)

type question969 struct {
	para969
	ans969
}

// para 是参数
// one 代表第一个参数
type para969 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans969 struct {
	one []int
}

func Benchmark_Problem969(b *testing.B) {

	qs := []question969{

		{
			para969{[]int{}},
			ans969{[]int{}},
		},

		{
			para969{[]int{1}},
			ans969{[]int{1}},
		},

		{
			para969{[]int{3, 2, 4, 1}},
			ans969{[]int{3, 4, 2, 3, 1, 2}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans969, q.para969
		(pancakeSort(p.one))
	}
}}}

package leetcode

import (
	"testing"
)

type question922 struct {
	para922
	ans922
}

// para 是参数
// one 代表第一个参数
type para922 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans922 struct {
	one []int
}

func Benchmark_Problem922(b *testing.B) {

	qs := []question922{

		{
			para922{[]int{}},
			ans922{[]int{}},
		},

		{
			para922{[]int{1}},
			ans922{[]int{}},
		},

		{
			para922{[]int{4, 2, 5, 7}},
			ans922{[]int{4, 5, 2, 7}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans922, q.para922
		(sortArrayByParityII(p.one))
	}
}}}

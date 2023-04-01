package leetcode

import (
	"testing"
)

type question455 struct {
	para455
	ans455
}

// para 是参数
// one 代表第一个参数
type para455 struct {
	g []int
	s []int
}

// ans 是答案
// one 代表第一个答案
type ans455 struct {
	one int
}

func Benchmark_Problem455(b *testing.B) {

	qs := []question455{

		{
			para455{[]int{1, 2, 3}, []int{1, 1}},
			ans455{1},
		},

		{
			para455{[]int{1, 2}, []int{1, 2, 3}},
			ans455{2},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans455, q.para455
		(findContentChildren(p.g, p.s))
	}
}}}

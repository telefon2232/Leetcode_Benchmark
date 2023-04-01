package leetcode

import (
	"testing"
)

type question135 struct {
	para135
	ans135
}

// para 是参数
// one 代表第一个参数
type para135 struct {
	ratings []int
}

// ans 是答案
// one 代表第一个答案
type ans135 struct {
	one int
}

func Benchmark_Problem135(b *testing.B) {

	qs := []question135{

		{
			para135{[]int{1, 0, 2}},
			ans135{5},
		},

		{
			para135{[]int{1, 2, 2}},
			ans135{4},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans135, q.para135
		(candy(p.ratings))
	}
}}}

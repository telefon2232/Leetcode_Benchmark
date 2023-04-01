package leetcode

import (
	"testing"
)

type question416 struct {
	para416
	ans416
}

// para 是参数
// one 代表第一个参数
type para416 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans416 struct {
	one bool
}

func Benchmark_Problem416(b *testing.B) {

	qs := []question416{

		{
			para416{[]int{1, 5, 11, 5}},
			ans416{true},
		},

		{
			para416{[]int{1, 2, 3, 5}},
			ans416{false},
		},

		{
			para416{[]int{1, 2, 5}},
			ans416{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans416, q.para416
		(canPartition(p.one))
	}
}}}

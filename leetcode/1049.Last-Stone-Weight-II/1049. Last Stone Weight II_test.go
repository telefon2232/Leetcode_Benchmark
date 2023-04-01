package leetcode

import (
	"testing"
)

type question1049 struct {
	para1049
	ans1049
}

// para 是参数
// one 代表第一个参数
type para1049 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans1049 struct {
	one int
}

func Benchmark_Problem1049(b *testing.B) {

	qs := []question1049{

		{
			para1049{[]int{2, 7, 4, 1, 8, 1}},
			ans1049{1},
		},

		{
			para1049{[]int{21, 26, 31, 33, 40}},
			ans1049{5},
		},

		{
			para1049{[]int{1, 2}},
			ans1049{1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans1049, q.para1049
		(lastStoneWeightII(p.one))
	}
}}}

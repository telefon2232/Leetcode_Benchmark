package leetcode

import (
	"testing"
)

type question53 struct {
	para53
	ans53
}

// para 是参数
// one 代表第一个参数
type para53 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans53 struct {
	one int
}

func Benchmark_Problem53(b *testing.B) {

	qs := []question53{

		{
			para53{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			ans53{6},
		},
		{
			para53{[]int{2, 7, 9, 3, 1}},
			ans53{22},
		},

		{
			para53{[]int{2}},
			ans53{2},
		},

		{
			para53{[]int{-1, -2}},
			ans53{-1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans53, q.para53
		(maxSubArray(p.one))
	}
}}}

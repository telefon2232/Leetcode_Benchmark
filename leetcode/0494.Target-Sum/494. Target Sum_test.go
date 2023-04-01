package leetcode

import (
	"testing"
)

type question494 struct {
	para494
	ans494
}

// para 是参数
// one 代表第一个参数
type para494 struct {
	nums []int
	S    int
}

// ans 是答案
// one 代表第一个答案
type ans494 struct {
	one int
}

func Benchmark_Problem494(b *testing.B) {

	qs := []question494{

		{
			para494{[]int{1, 1, 1, 1, 1}, 3},
			ans494{5},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans494, q.para494
		(findTargetSumWays(p.nums, p.S))
	}
}}}

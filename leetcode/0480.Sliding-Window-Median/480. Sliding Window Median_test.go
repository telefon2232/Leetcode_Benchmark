package leetcode

import (
	"testing"
)

type question480 struct {
	para480
	ans480
}

// para 是参数
// one 代表第一个参数
type para480 struct {
	one []int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans480 struct {
	one []int
}

func Benchmark_Problem480(b *testing.B) {

	qs := []question480{

		{
			para480{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3},
			ans480{[]int{1, -1, -1, 3, 5, 6}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans480, q.para480
		(medianSlidingWindow1(p.one, p.k))
	}
}}}

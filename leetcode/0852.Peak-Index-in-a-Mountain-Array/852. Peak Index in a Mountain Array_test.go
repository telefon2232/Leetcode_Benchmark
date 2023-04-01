package leetcode

import (
	"testing"
)

type question852 struct {
	para852
	ans852
}

// para 是参数
// one 代表第一个参数
type para852 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans852 struct {
	one int
}

func Benchmark_Problem852(b *testing.B) {

	qs := []question852{

		{
			para852{[]int{0, 1, 0}},
			ans852{1},
		},

		{
			para852{[]int{0, 2, 1, 0}},
			ans852{1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans852, q.para852
		(peakIndexInMountainArray(p.one))
	}
}}}

package leetcode

import (
	"testing"
)

type question274 struct {
	para274
	ans274
}

// para 是参数
// one 代表第一个参数
type para274 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans274 struct {
	one int
}

func Benchmark_Problem274(b *testing.B) {

	qs := []question274{

		{
			para274{[]int{3, 6, 9, 1}},
			ans274{3},
		},
		{
			para274{[]int{1}},
			ans274{1},
		},

		{
			para274{[]int{}},
			ans274{0},
		},

		{
			para274{[]int{3, 0, 6, 1, 5}},
			ans274{3},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans274, q.para274
		(hIndex(p.one))
	}
}}}

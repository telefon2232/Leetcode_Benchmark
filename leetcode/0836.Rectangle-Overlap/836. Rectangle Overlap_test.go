package leetcode

import (
	"testing"
)

type question836 struct {
	para836
	ans836
}

// para 是参数
// one 代表第一个参数
type para836 struct {
	rec1 []int
	rec2 []int
}

// ans 是答案
// one 代表第一个答案
type ans836 struct {
	one bool
}

func Benchmark_Problem836(b *testing.B) {

	qs := []question836{

		{
			para836{[]int{0, 0, 2, 2}, []int{1, 1, 3, 3}},
			ans836{true},
		},

		{
			para836{[]int{0, 0, 1, 1}, []int{1, 0, 2, 1}},
			ans836{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans836, q.para836
		(isRectangleOverlap(p.rec1, p.rec2))
	}
}}}

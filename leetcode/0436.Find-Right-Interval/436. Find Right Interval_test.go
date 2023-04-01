package leetcode

import (
	"testing"
)

type question436 struct {
	para436
	ans436
}

// para 是参数
// one 代表第一个参数
type para436 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans436 struct {
	one []int
}

func Benchmark_Problem436(b *testing.B) {

	qs := []question436{

		{
			para436{[][]int{{3, 4}, {2, 3}, {1, 2}}},
			ans436{[]int{-1, 0, 1}},
		},

		{
			para436{[][]int{{1, 4}, {2, 3}, {3, 4}}},
			ans436{[]int{-1, 2, -1}},
		},

		{
			para436{[][]int{{1, 2}}},
			ans436{[]int{-1}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans436, q.para436
		(findRightInterval(p.one))
	}
}}}

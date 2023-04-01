package leetcode

import (
	"testing"
)

type question904 struct {
	para904
	ans904
}

// para 是参数
// one 代表第一个参数
type para904 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans904 struct {
	one int
}

func Benchmark_Problem904(b *testing.B) {

	qs := []question904{

		{
			para904{[]int{1, 1}},
			ans904{2},
		},

		{
			para904{[]int{1, 2, 1}},
			ans904{3},
		},
		{
			para904{[]int{0, 1, 2, 2}},
			ans904{3},
		},

		{
			para904{[]int{1, 2, 3, 2, 2}},
			ans904{4},
		},

		{
			para904{[]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}},
			ans904{5},
		},

		{
			para904{[]int{4, 5}},
			ans904{2},
		},

		{
			para904{[]int{1}},
			ans904{1},
		},

		{
			para904{[]int{}},
			ans904{},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans904, q.para904
		(totalFruit(p.one))
	}
}}}

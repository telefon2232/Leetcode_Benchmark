package leetcode

import (
	"testing"
)

type question475 struct {
	para475
	ans475
}

// para 是参数
// one 代表第一个参数
type para475 struct {
	houses  []int
	heaters []int
}

// ans 是答案
// one 代表第一个答案
type ans475 struct {
	one int
}

func Benchmark_Problem475(b *testing.B) {

	qs := []question475{

		{
			para475{[]int{1, 2, 3}, []int{2}},
			ans475{1},
		},

		{
			para475{[]int{1, 2, 3, 4}, []int{1, 4}},
			ans475{1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans475, q.para475
		(findRadius(p.houses, p.heaters))
	}
}}}

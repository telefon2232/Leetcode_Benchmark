package leetcode

import (
	"testing"
)

type question198 struct {
	para198
	ans198
}

// para 是参数
// one 代表第一个参数
type para198 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans198 struct {
	one int
}

func Benchmark_Problem198(b *testing.B) {

	qs := []question198{

		{
			para198{[]int{1, 2}},
			ans198{2},
		},

		{
			para198{[]int{1, 2, 3, 1}},
			ans198{4},
		},
		{
			para198{[]int{2, 7, 9, 3, 1}},
			ans198{12},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans198, q.para198
		(rob198(p.one))
	}
}}}

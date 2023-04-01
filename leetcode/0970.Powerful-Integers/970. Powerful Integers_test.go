package leetcode

import (
	"testing"
)

type question970 struct {
	para970
	ans970
}

// para 是参数
// one 代表第一个参数
type para970 struct {
	one int
	two int
	b   int
}

// ans 是答案
// one 代表第一个答案
type ans970 struct {
	one []int
}

func Benchmark_Problem970(b *testing.B) {

	qs := []question970{

		{
			para970{2, 3, 10},
			ans970{[]int{2, 3, 4, 5, 7, 9, 10}},
		},

		{
			para970{3, 5, 15},
			ans970{[]int{2, 4, 6, 8, 10, 14}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans970, q.para970
		(powerfulIntegers(p.one, p.two, p.b))
	}
}}}

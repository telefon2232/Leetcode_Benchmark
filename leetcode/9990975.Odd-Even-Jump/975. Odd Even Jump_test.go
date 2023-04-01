package leetcode

import (
	"testing"
)

type question975 struct {
	para975
	ans975
}

// para 是参数
// one 代表第一个参数
type para975 struct {
	A []int
}

// ans 是答案
// one 代表第一个答案
type ans975 struct {
	one int
}

func Benchmark_Problem975(b *testing.B) {

	qs := []question975{

		{
			para975{[]int{10, 13, 12, 14, 15}},
			ans975{2},
		},

		{
			para975{[]int{2, 3, 1, 1, 4}},
			ans975{3},
		},

		{
			para975{[]int{5, 1, 3, 4, 2}},
			ans975{3},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans975, q.para975
		(oddEvenJumps(p.A))
	}
}}}

package leetcode

import (
	"testing"
)

type question961 struct {
	para961
	ans961
}

// para 是参数
// one 代表第一个参数
type para961 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans961 struct {
	one int
}

func Benchmark_Problem961(b *testing.B) {

	qs := []question961{
		{
			para961{[]int{1, 2, 3, 3}},
			ans961{3},
		},

		{
			para961{[]int{2, 1, 2, 5, 3, 2}},
			ans961{2},
		},

		{
			para961{[]int{5, 1, 5, 2, 5, 3, 5, 4}},
			ans961{5},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans961, q.para961
		(repeatedNTimes(p.one))
	}
}}}

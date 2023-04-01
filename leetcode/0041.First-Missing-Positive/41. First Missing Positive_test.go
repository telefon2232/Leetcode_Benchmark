package leetcode

import (
	"testing"
)

type question41 struct {
	para41
	ans41
}

// para 是参数
// one 代表第一个参数
type para41 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans41 struct {
	one int
}

func Benchmark_Problem41(b *testing.B) {

	qs := []question41{

		{
			para41{[]int{10, -1, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5, -3}},
			ans41{6},
		},

		{
			para41{[]int{10, -1, 8, 6, 7, 3, -2, 5, 4, 2, 1, -3}},
			ans41{9},
		},

		{
			para41{[]int{1}},
			ans41{2},
		},

		{
			para41{[]int{0, 2, 2, 1, 1}},
			ans41{3},
		},

		{
			para41{[]int{}},
			ans41{1},
		},

		{
			para41{[]int{1, 2, 0}},
			ans41{3},
		},

		{
			para41{[]int{3, 4, -1, 1}},
			ans41{2},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans41, q.para41
		(firstMissingPositive(p.one))
	}
}}}

package leetcode

import (
	"testing"
)

type question665 struct {
	para665
	ans665
}

// para 是参数
// one 代表第一个参数
type para665 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans665 struct {
	one bool
}

func Benchmark_Problem665(b *testing.B) {

	qs := []question665{

		{
			para665{[]int{4, 2, 3}},
			ans665{true},
		},

		{
			para665{[]int{4, 2, 1}},
			ans665{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans665, q.para665
				(checkPossibility(p.nums))
			}
		}
	}
}

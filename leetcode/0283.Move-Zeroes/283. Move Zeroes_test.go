package leetcode

import (
	"testing"
)

type question283 struct {
	para283
	ans283
}

// para 是参数
// one 代表第一个参数
type para283 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans283 struct {
	one []int
}

func Benchmark_Problem283(b *testing.B) {

	qs := []question283{

		{
			para283{[]int{1, 0, 1}},
			ans283{[]int{1, 1, 0}},
		},

		{
			para283{[]int{0, 1, 0, 3, 0, 12}},
			ans283{[]int{1, 3, 12, 0, 0, 0}},
		},

		{
			para283{[]int{0, 1, 0, 3, 0, 0, 0, 0, 1, 12}},
			ans283{[]int{1, 3, 1, 12, 0, 0, 0, 0, 0}},
		},

		{
			para283{[]int{0, 0, 0, 0, 0, 0, 0, 0, 12, 1}},
			ans283{[]int{12, 1, 0, 0, 0, 0, 0, 0, 0, 0}},
		},

		{
			para283{[]int{0, 0, 0, 0, 0}},
			ans283{[]int{0, 0, 0, 0, 0}},
		},

		{
			para283{[]int{1}},
			ans283{[]int{1}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans283, q.para283

				moveZeroes(p.one)

			}
		}
	}
}

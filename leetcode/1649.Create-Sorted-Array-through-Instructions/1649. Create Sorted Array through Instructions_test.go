package leetcode

import (
	"testing"
)

type question1649 struct {
	para1649
	ans1649
}

// para 是参数
// one 代表第一个参数
type para1649 struct {
	instructions []int
}

// ans 是答案
// one 代表第一个答案
type ans1649 struct {
	one int
}

func Benchmark_Problem1649(b *testing.B) {

	qs := []question1649{

		{
			para1649{[]int{1, 5, 6, 2}},
			ans1649{1},
		},

		{
			para1649{[]int{1, 2, 3, 6, 5, 4}},
			ans1649{3},
		},

		{
			para1649{[]int{1, 3, 3, 3, 2, 4, 2, 1, 2}},
			ans1649{4},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1649, q.para1649
				(createSortedArray(p.instructions))
			}
		}
	}
}

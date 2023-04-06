package leetcode

import (
	"testing"
)

type question506 struct {
	para506
	ans506
}

// para 是参数
type para506 struct {
	score []int
}

// ans 是答案
type ans506 struct {
	ans []string
}

func Benchmark_Problem506(b *testing.B) {

	qs := []question506{

		{
			para506{[]int{5, 4, 3, 2, 1}},
			ans506{[]string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}},
		},

		{
			para506{[]int{10, 3, 8, 9, 4}},
			ans506{[]string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans506, q.para506
				(findRelativeRanks(p.score))
			}
		}
	}
}

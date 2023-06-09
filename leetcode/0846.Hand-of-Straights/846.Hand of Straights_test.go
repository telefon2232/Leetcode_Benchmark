package leetcode

import (
	"testing"
)

type question846 struct {
	para846
	ans846
}

// para 是参数
type para846 struct {
	hand      []int
	groupSize int
}

// ans 是答案
type ans846 struct {
	ans bool
}

func Benchmark_Problem846(b *testing.B) {

	qs := []question846{

		{
			para846{[]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3},
			ans846{true},
		},

		{
			para846{[]int{1, 2, 3, 4, 5}, 4},
			ans846{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans846, q.para846
				(isNStraightHand(p.hand, p.groupSize))
			}
		}
	}
}

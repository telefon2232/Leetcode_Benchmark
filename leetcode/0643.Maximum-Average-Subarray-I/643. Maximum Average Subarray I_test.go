package leetcode

import (
	"testing"
)

type question643 struct {
	para643
	ans643
}

// para 是参数
// one 代表第一个参数
type para643 struct {
	nums []int
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans643 struct {
	one float64
}

func Benchmark_Problem643(b *testing.B) {

	qs := []question643{

		{
			para643{[]int{1, 12, -5, -6, 50, 3}, 4},
			ans643{12.75},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans643, q.para643
				(findMaxAverage(p.nums, p.k))
			}
		}
	}
}

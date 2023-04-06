package leetcode

import (
	"testing"
)

type question495 struct {
	para495
	ans495
}

// para 是参数
type para495 struct {
	timeSeries []int
	duration   int
}

// ans 是答案
type ans495 struct {
	ans int
}

func Benchmark_Problem495(b *testing.B) {

	qs := []question495{

		{
			para495{[]int{1, 4}, 2},
			ans495{4},
		},

		{
			para495{[]int{1, 2}, 2},
			ans495{3},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans495, q.para495
				(findPoisonedDuration(p.timeSeries, p.duration))
			}
		}
	}
}

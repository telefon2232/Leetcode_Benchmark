package leetcode

import (
	"testing"
)

type question477 struct {
	para477
	ans477
}

// para 是参数
// one 代表第一个参数
type para477 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans477 struct {
	one int
}

func Benchmark_Problem477(b *testing.B) {

	qs := []question477{

		{
			para477{[]int{4, 14, 2}},
			ans477{6},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans477, q.para477
				(totalHammingDistance(p.one))
			}
		}
	}
}

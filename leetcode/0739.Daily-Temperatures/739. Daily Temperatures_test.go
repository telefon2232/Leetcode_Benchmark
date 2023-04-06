package leetcode

import (
	"testing"
)

type question739 struct {
	para739
	ans739
}

// para 是参数
// one 代表第一个参数
type para739 struct {
	s []int
}

// ans 是答案
// one 代表第一个答案
type ans739 struct {
	one []int
}

func Benchmark_Problem739(b *testing.B) {

	qs := []question739{

		{
			para739{[]int{73, 74, 75, 71, 69, 72, 76, 73}},
			ans739{[]int{1, 1, 4, 2, 1, 1, 0, 0}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans739, q.para739
				(dailyTemperatures(p.s))
			}
		}
	}
}

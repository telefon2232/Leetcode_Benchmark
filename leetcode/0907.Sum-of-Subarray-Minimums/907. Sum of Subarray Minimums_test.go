package leetcode

import (
	"testing"
)

type question907 struct {
	para907
	ans907
}

// para 是参数
// one 代表第一个参数
type para907 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans907 struct {
	one int
}

func Benchmark_Problem907(b *testing.B) {

	qs := []question907{

		{
			para907{[]int{3, 1, 2, 4}},
			ans907{17},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans907, q.para907
				(sumSubarrayMins(p.one))
			}
		}
	}
}

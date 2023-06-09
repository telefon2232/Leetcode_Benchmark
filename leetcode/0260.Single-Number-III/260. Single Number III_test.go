package leetcode

import (
	"testing"
)

type question260 struct {
	para260
	ans260
}

// para 是参数
// one 代表第一个参数
type para260 struct {
	s []int
}

// ans 是答案
// one 代表第一个答案
type ans260 struct {
	one []int
}

func Benchmark_Problem260(b *testing.B) {

	qs := []question260{

		{
			para260{[]int{1, 2, 1, 3, 2, 5}},
			ans260{[]int{3, 5}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans260, q.para260
				(singleNumberIII(p.s))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question877 struct {
	para877
	ans877
}

// para 是参数
// one 代表第一个参数
type para877 struct {
	piles []int
}

// ans 是答案
// one 代表第一个答案
type ans877 struct {
	one bool
}

func Benchmark_Problem877(b *testing.B) {

	qs := []question877{

		{
			para877{[]int{5, 3, 4, 5}},
			ans877{true},
		},
	}

	for _, q := range qs {
		_, p := q.ans877, q.para877
		(stoneGame(p.piles))
	}
}

package leetcode

import (
	"testing"
)

type question374 struct {
	para374
	ans374
}

// para 是参数
// one 代表第一个参数
type para374 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans374 struct {
	one int
}

func Benchmark_Problem374(b *testing.B) {

	qs := []question374{

		{
			para374{10},
			ans374{6},
		},

		{
			para374{1},
			ans374{1},
		},

		{
			para374{2},
			ans374{1},
		},

		{
			para374{2},
			ans374{2},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans374, q.para374
				(guessNumber(p.n))
			}
		}
	}
}

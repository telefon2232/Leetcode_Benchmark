package leetcode

import (
	"testing"
)

type question278 struct {
	para278
	ans278
}

// para 是参数
// one 代表第一个参数
type para278 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans278 struct {
	one int
}

func Benchmark_Problem278(b *testing.B) {

	qs := []question278{

		{
			para278{5},
			ans278{4},
		},
		{
			para278{1},
			ans278{1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans278, q.para278
				(firstBadVersion(p.n))
			}
		}
	}
}

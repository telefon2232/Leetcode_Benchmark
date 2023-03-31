package leetcode

import (
	"testing"
)

type question441 struct {
	para441
	ans441
}

// para 是参数
// one 代表第一个参数
type para441 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans441 struct {
	one int
}

func Benchmark_Problem441(b *testing.B) {

	qs := []question441{

		{
			para441{5},
			ans441{2},
		},

		{
			para441{8},
			ans441{3},
		},
	}

	for _, q := range qs {
		_, p := q.ans441, q.para441
		(arrangeCoins(p.n))
	}
}

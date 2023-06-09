package leetcode

import (
	"testing"
)

type question793 struct {
	para793
	ans793
}

// para 是参数
// one 代表第一个参数
type para793 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans793 struct {
	one int
}

func Benchmark_Problem793(b *testing.B) {

	qs := []question793{

		{
			para793{0},
			ans793{5},
		},

		{
			para793{5},
			ans793{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans793, q.para793
				(preimageSizeFZF(p.one))
			}
		}
	}
}

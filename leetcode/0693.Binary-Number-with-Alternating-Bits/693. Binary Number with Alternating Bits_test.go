package leetcode

import (
	"testing"
)

type question693 struct {
	para693
	ans693
}

// para 是参数
// one 代表第一个参数
type para693 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans693 struct {
	one bool
}

func Benchmark_Problem693(b *testing.B) {

	qs := []question693{

		{
			para693{5},
			ans693{true},
		},

		{
			para693{7},
			ans693{false},
		},

		{
			para693{11},
			ans693{false},
		},

		{
			para693{10},
			ans693{true},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans693, q.para693
				(hasAlternatingBits(p.one))
			}
		}
	}
}

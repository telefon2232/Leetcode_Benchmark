package leetcode

import (
	"testing"
)

type question995 struct {
	para995
	ans995
}

// para 是参数
// one 代表第一个参数
type para995 struct {
	one []int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans995 struct {
	one int
}

func Benchmark_Problem995(b *testing.B) {

	qs := []question995{

		{
			para995{[]int{0, 1, 0}, 1},
			ans995{2},
		},

		{
			para995{[]int{1, 1, 0}, 2},
			ans995{-1},
		},

		{
			para995{[]int{0, 0, 0, 1, 0, 1, 1, 0}, 3},
			ans995{3},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans995, q.para995
				(minKBitFlips(p.one, p.k))
			}
		}
	}
}

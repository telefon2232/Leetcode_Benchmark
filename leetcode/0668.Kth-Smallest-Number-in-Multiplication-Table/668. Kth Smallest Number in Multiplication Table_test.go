package leetcode

import (
	"testing"
)

type question668 struct {
	para668
	ans668
}

// para 是参数
// one 代表第一个参数
type para668 struct {
	m int
	n int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans668 struct {
	one int
}

func Benchmark_Problem668(b *testing.B) {

	qs := []question668{

		{
			para668{3, 3, 5},
			ans668{3},
		},

		{
			para668{2, 3, 6},
			ans668{6},
		},

		{
			para668{1, 3, 2},
			ans668{2},
		},

		{
			para668{42, 34, 401},
			ans668{126},
		},

		{
			para668{7341, 13535, 12330027},
			ans668{2673783},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans668, q.para668
				(findKthNumber(p.m, p.n, p.k))
			}
		}
	}
}

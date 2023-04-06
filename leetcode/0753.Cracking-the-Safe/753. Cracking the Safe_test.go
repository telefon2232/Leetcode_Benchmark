package leetcode

import (
	"testing"
)

type question753 struct {
	para753
	ans753
}

// para 是参数
// one 代表第一个参数
type para753 struct {
	n int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans753 struct {
	one string
}

func Benchmark_Problem753(b *testing.B) {

	qs := []question753{

		{
			para753{1, 2},
			ans753{"01"},
		},

		{
			para753{2, 2},
			ans753{"00110"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans753, q.para753
				(crackSafe(p.n, p.k))
			}
		}
	}
}

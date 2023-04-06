package leetcode

import (
	"testing"
)

type question819 struct {
	para819
	ans819
}

// para 是参数
// one 代表第一个参数
type para819 struct {
	one string
	b   []string
}

// ans 是答案
// one 代表第一个答案
type ans819 struct {
	one string
}

func Benchmark_Problem819(b *testing.B) {

	qs := []question819{

		{
			para819{"Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}},
			ans819{"ball"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans819, q.para819
				(mostCommonWord(p.one, p.b))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question792 struct {
	para792
	ans792
}

// para 是参数
// one 代表第一个参数
type para792 struct {
	s     string
	words []string
}

// ans 是答案
// one 代表第一个答案
type ans792 struct {
	one int
}

func Benchmark_Problem792(b *testing.B) {

	qs := []question792{

		{
			para792{"abcde", []string{"a", "bb", "acd", "ace"}},
			ans792{3},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans792, q.para792
				(numMatchingSubseq(p.s, p.words))
			}
		}
	}
}

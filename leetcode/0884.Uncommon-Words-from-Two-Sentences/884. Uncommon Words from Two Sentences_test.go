package leetcode

import (
	"testing"
)

type question884 struct {
	para884
	ans884
}

// para 是参数
// one 代表第一个参数
type para884 struct {
	A string
	B string
}

// ans 是答案
// one 代表第一个答案
type ans884 struct {
	one []string
}

func Benchmark_Problem884(b *testing.B) {

	qs := []question884{

		{
			para884{"this apple is sweet", "this apple is sour"},
			ans884{[]string{"sweet", "sour"}},
		},

		{
			para884{"apple apple", "banana"},
			ans884{[]string{"banana"}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans884, q.para884
				(uncommonFromSentences(p.A, p.B))
			}
		}
	}
}

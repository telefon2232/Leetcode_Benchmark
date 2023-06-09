package leetcode

import (
	"testing"
)

type question345 struct {
	para345
	ans345
}

// para 是参数
// one 代表第一个参数
type para345 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans345 struct {
	one string
}

func Benchmark_Problem345(b *testing.B) {

	qs := []question345{

		{
			para345{"hello"},
			ans345{"holle"},
		},

		{
			para345{"leetcode"},
			ans345{"leotcede"},
		},

		{
			para345{"aA"},
			ans345{"Aa"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans345, q.para345
				(reverseVowels(p.one))
			}
		}
	}
}

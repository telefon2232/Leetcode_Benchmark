package leetcode

import (
	"testing"
)

type question205 struct {
	para205
	ans205
}

// para 是参数
// one 代表第一个参数
type para205 struct {
	one string
	two string
}

// ans 是答案
// one 代表第一个答案
type ans205 struct {
	one bool
}

func Benchmark_Problem205(b *testing.B) {

	qs := []question205{

		{
			para205{"egg", "add"},
			ans205{true},
		},

		{
			para205{"foo", "bar"},
			ans205{false},
		},

		{
			para205{"paper", "title"},
			ans205{true},
		},

		{
			para205{"", ""},
			ans205{true},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans205, q.para205
				(isIsomorphic(p.one, p.two))
			}
		}
	}
}

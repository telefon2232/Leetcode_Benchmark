package leetcode

import (
	"testing"
)

type question383 struct {
	para383
	ans383
}

// para 是参数
type para383 struct {
	ransomNote string
	magazine   string
}

// ans 是答案
type ans383 struct {
	ans bool
}

func Benchmark_Problem383(b *testing.B) {

	qs := []question383{

		{
			para383{"a", "b"},
			ans383{false},
		},

		{
			para383{"aa", "ab"},
			ans383{false},
		},

		{
			para383{"aa", "aab"},
			ans383{true},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans383, q.para383
				(canConstruct(p.ransomNote, p.magazine))
			}
		}
	}
}

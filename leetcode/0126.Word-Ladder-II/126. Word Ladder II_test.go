package leetcode

import (
	"testing"
)

type question126 struct {
	para126
	ans126
}

// para 是参数
// one 代表第一个参数
type para126 struct {
	b string
	e string
	w []string
}

// ans 是答案
// one 代表第一个答案
type ans126 struct {
	one [][]string
}

func Benchmark_Problem126(b *testing.B) {

	qs := []question126{
		{
			para126{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}},
			ans126{[][]string{{"hit", "hot", "dot", "dog", "cog"}, {"hit", "hot", "lot", "log", "cog"}}},
		},

		{
			para126{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}},
			ans126{[][]string{}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans126, q.para126
				(findLadders(p.b, p.e, p.w))
			}
		}
	}
}

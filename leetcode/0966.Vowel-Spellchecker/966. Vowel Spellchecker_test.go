package leetcode

import (
	"testing"
)

type question966 struct {
	para966
	ans966
}

// para 是参数
// one 代表第一个参数
type para966 struct {
	wordlist []string
	queries  []string
}

// ans 是答案
// one 代表第一个答案
type ans966 struct {
	one []string
}

func Benchmark_Problem966(b *testing.B) {

	qs := []question966{
		{
			para966{[]string{"KiTe", "kite", "hare", "Hare"}, []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}},
			ans966{[]string{"kite", "KiTe", "KiTe", "Hare", "hare", "", "", "KiTe", "", "KiTe"}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans966, q.para966
				(spellchecker(p.wordlist, p.queries))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question433 struct {
	para433
	ans433
}

// para 是参数
// one 代表第一个参数
type para433 struct {
	start string
	end   string
	bank  []string
}

// ans 是答案
// one 代表第一个答案
type ans433 struct {
	one int
}

func Benchmark_Problem433(b *testing.B) {

	qs := []question433{
		{
			para433{"AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}},
			ans433{1},
		},

		{
			para433{"AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}},
			ans433{2},
		},

		{
			para433{"AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}},
			ans433{3},
		},

		{
			para433{"AACCGGTT", "AACCGGTA", []string{}},
			ans433{-1},
		},
	}

	for _, q := range qs {
		_, p := q.ans433, q.para433
		(minMutation(p.start, p.end, p.bank))
	}
}

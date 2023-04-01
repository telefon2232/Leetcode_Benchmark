package leetcode

import (
	"testing"
)

type question451 struct {
	para451
	ans451
}

// para 是参数
// one 代表第一个参数
type para451 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans451 struct {
	one string
}

func Benchmark_Problem451(b *testing.B) {

	qs := []question451{
		{
			para451{"tree"},
			ans451{"eert"},
		},

		{
			para451{"cccaaa"},
			ans451{"cccaaa"},
		},

		{
			para451{"Aabb"},
			ans451{"bbAa"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans451, q.para451
		(frequencySort(p.one))
	}
}}}

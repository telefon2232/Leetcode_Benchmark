package leetcode

import (
	"testing"
)

type question720 struct {
	para720
	ans720
}

// para 是参数
// one 代表第一个参数
type para720 struct {
	w []string
}

// ans 是答案
// one 代表第一个答案
type ans720 struct {
	one string
}

func Benchmark_Problem720(b *testing.B) {

	qs := []question720{

		{
			para720{[]string{"w", "wo", "wor", "worl", "world"}},
			ans720{"world"},
		},

		{
			para720{[]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}},
			ans720{"apple"},
		},
	}

	for _, q := range qs {
		_, p := q.ans720, q.para720
		(longestWord(p.w))
	}
}

package leetcode

import (
	"testing"
)

type question395 struct {
	para395
	ans395
}

// para 是参数
// one 代表第一个参数
type para395 struct {
	s string
	k int
}

// ans 是答案
// one 代表第一个答案
type ans395 struct {
	one int
}

func Benchmark_Problem395(b *testing.B) {

	qs := []question395{

		{
			para395{"aaabb", 3},
			ans395{3},
		},

		{
			para395{"ababbc", 2},
			ans395{5},
		},
	}

	for _, q := range qs {
		_, p := q.ans395, q.para395
		(longestSubstring(p.s, p.k))
	}
}

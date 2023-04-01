package leetcode

import (
	"testing"
)

type question402 struct {
	para402
	ans402
}

// para 是参数
// one 代表第一个参数
type para402 struct {
	num string
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans402 struct {
	one string
}

func Benchmark_Problem402(b *testing.B) {

	qs := []question402{
		{
			para402{"10", 1},
			ans402{"0"},
		},

		{
			para402{"1111111", 3},
			ans402{"1111"},
		},

		{
			para402{"5337", 2},
			ans402{"33"},
		},

		{
			para402{"112", 1},
			ans402{"11"},
		},

		{
			para402{"1432219", 3},
			ans402{"1219"},
		},

		{
			para402{"10200", 1},
			ans402{"200"},
		},

		{
			para402{"10", 2},
			ans402{"0"},
		},

		{
			para402{"19", 2},
			ans402{"0"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans402, q.para402
		(removeKdigits(p.num, p.k))
	}
}}}

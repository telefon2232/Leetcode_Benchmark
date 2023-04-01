package leetcode

import (
	"testing"
)

type question887 struct {
	para887
	ans887
}

// para 是参数
// one 代表第一个参数
type para887 struct {
	k int
	n int
}

// ans 是答案
// one 代表第一个答案
type ans887 struct {
	one int
}

func Benchmark_Problem887(b *testing.B) {

	qs := []question887{

		{
			para887{1, 2},
			ans887{2},
		},

		{
			para887{2, 6},
			ans887{3},
		},

		{
			para887{3, 14},
			ans887{4},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans887, q.para887
		(superEggDrop(p.k, p.n))
	}
}}}

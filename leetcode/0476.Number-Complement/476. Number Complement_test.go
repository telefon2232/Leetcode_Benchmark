package leetcode

import (
	"testing"
)

type question476 struct {
	para476
	ans476
}

// para 是参数
// one 代表第一个参数
type para476 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans476 struct {
	one int
}

func Benchmark_Problem476(b *testing.B) {

	qs := []question476{

		{
			para476{5},
			ans476{2},
		},

		{
			para476{1},
			ans476{0},
		},
	}

	for _, q := range qs {
		_, p := q.ans476, q.para476
		(findComplement(p.one))
	}
}

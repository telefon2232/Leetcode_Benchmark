package leetcode

import (
	"testing"
)

type question258 struct {
	para258
	ans258
}

// para 是参数
// one 代表第一个参数
type para258 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans258 struct {
	one int
}

func Benchmark_Problem258(b *testing.B) {

	qs := []question258{

		{
			para258{38},
			ans258{2},
		},

		{
			para258{88},
			ans258{7},
		},

		{
			para258{96},
			ans258{6},
		},
	}

	for _, q := range qs {
		_, p := q.ans258, q.para258
		(addDigits(p.one))
	}
}

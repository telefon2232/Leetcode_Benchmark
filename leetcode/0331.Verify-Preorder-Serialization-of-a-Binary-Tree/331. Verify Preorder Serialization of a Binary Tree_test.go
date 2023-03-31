package leetcode

import (
	"testing"
)

type question331 struct {
	para331
	ans331
}

// para 是参数
// one 代表第一个参数
type para331 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans331 struct {
	one bool
}

func Benchmark_Problem331(b *testing.B) {

	qs := []question331{

		{
			para331{"9,3,4,#,#,1,#,#,2,#,6,#,#"},
			ans331{true},
		},

		{
			para331{"1,#"},
			ans331{false},
		},

		{
			para331{"9,#,#,1"},
			ans331{false},
		},
	}

	for _, q := range qs {
		_, p := q.ans331, q.para331
		(isValidSerialization(p.one))
	}
}

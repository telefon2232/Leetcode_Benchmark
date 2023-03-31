package leetcode

import (
	"testing"
)

type question1017 struct {
	para1017
	ans1017
}

// para 是参数
// one 代表第一个参数
type para1017 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans1017 struct {
	one string
}

func Benchmark_Problem1017(b *testing.B) {

	qs := []question1017{

		{
			para1017{2},
			ans1017{"110"},
		},

		{
			para1017{3},
			ans1017{"111"},
		},

		{
			para1017{4},
			ans1017{"110"},
		},
	}

	for _, q := range qs {
		_, p := q.ans1017, q.para1017
		(baseNeg2(p.one))
	}
}

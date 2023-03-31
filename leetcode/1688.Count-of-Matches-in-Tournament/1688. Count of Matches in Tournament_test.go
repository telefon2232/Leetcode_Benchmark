package leetcode

import (
	"testing"
)

type question1688 struct {
	para1688
	ans1688
}

// para 是参数
// one 代表第一个参数
type para1688 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans1688 struct {
	one int
}

func Benchmark_Problem1688(b *testing.B) {

	qs := []question1688{

		{
			para1688{7},
			ans1688{6},
		},

		{
			para1688{14},
			ans1688{13},
		},
	}

	for _, q := range qs {
		_, p := q.ans1688, q.para1688
		(numberOfMatches(p.n))
	}
}

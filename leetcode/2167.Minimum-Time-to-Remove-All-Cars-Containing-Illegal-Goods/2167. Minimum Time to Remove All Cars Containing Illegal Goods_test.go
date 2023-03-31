package leetcode

import (
	"testing"
)

type question2167 struct {
	para2167
	ans2167
}

// para 是参数
// one 代表第一个参数
type para2167 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans2167 struct {
	one int
}

func Benchmark_Problem2167(b *testing.B) {

	qs := []question2167{

		{
			para2167{"1100101"},
			ans2167{5},
		},

		{
			para2167{"0010"},
			ans2167{2},
		},

		{
			para2167{"1100111101"},
			ans2167{8},
		},

		{
			para2167{"1001010101"},
			ans2167{8},
		},
	}

	for _, q := range qs {
		_, p := q.ans2167, q.para2167
		(minimumTime(p.s))
	}
}

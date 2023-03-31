package leetcode

import (
	"testing"
)

type question869 struct {
	para869
	ans869
}

// para 是参数
// one 代表第一个参数
type para869 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans869 struct {
	one bool
}

func Benchmark_Problem869(b *testing.B) {

	qs := []question869{

		{
			para869{1},
			ans869{true},
		},

		{
			para869{10},
			ans869{false},
		},

		{
			para869{16},
			ans869{true},
		},

		{
			para869{24},
			ans869{false},
		},

		{
			para869{46},
			ans869{true},
		},

		{
			para869{100},
			ans869{false},
		},

		{
			para869{123453242},
			ans869{false},
		},
	}

	for _, q := range qs {
		_, p := q.ans869, q.para869
		(reorderedPowerOf2(p.n))
	}
}

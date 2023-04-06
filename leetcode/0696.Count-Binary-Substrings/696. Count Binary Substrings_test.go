package leetcode

import (
	"testing"
)

type question696 struct {
	para696
	ans696
}

// para 是参数
// one 代表第一个参数
type para696 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans696 struct {
	one int
}

func Benchmark_Problem696(b *testing.B) {

	qs := []question696{

		{
			para696{"00110011"},
			ans696{6},
		},

		{
			para696{"10101"},
			ans696{4},
		},

		{
			para696{"0110001111"},
			ans696{6},
		},

		{
			para696{"0001111"},
			ans696{3},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans696, q.para696
				(countBinarySubstrings(p.one))
			}
		}
	}
}

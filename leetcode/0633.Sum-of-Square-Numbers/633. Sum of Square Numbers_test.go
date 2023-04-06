package leetcode

import (
	"testing"
)

type question633 struct {
	para633
	ans633
}

// para 是参数
// one 代表第一个参数
type para633 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans633 struct {
	one bool
}

func Benchmark_Problem633(b *testing.B) {

	qs := []question633{

		{
			para633{1},
			ans633{true},
		},

		{
			para633{2},
			ans633{true},
		},

		{
			para633{3},
			ans633{false},
		},

		{
			para633{4},
			ans633{true},
		},

		{
			para633{5},
			ans633{true},
		},

		{
			para633{6},
			ans633{false},
		},

		{
			para633{104976},
			ans633{true},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans633, q.para633
				(judgeSquareSum(p.one))
			}
		}
	}
}

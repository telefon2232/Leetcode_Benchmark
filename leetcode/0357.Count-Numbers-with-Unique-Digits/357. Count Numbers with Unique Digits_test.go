package leetcode

import (
	"testing"
)

type question357 struct {
	para357
	ans357
}

// para 是参数
// one 代表第一个参数
type para357 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans357 struct {
	one int
}

func Benchmark_Problem357(b *testing.B) {

	qs := []question357{

		{
			para357{1},
			ans357{10},
		},

		{
			para357{2},
			ans357{91},
		},

		{
			para357{3},
			ans357{739},
		},

		{
			para357{4},
			ans357{5275},
		},

		{
			para357{5},
			ans357{32491},
		},

		{
			para357{6},
			ans357{168571},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans357, q.para357
				(countNumbersWithUniqueDigits(p.one))
			}
		}
	}
}

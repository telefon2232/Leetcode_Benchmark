package leetcode

import (
	"testing"
)

type question456 struct {
	para456
	ans456
}

// para 是参数
// one 代表第一个参数
type para456 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans456 struct {
	one bool
}

func Benchmark_Problem456(b *testing.B) {

	qs := []question456{

		{
			para456{[]int{}},
			ans456{false},
		},

		{
			para456{[]int{1, 2, 3, 4}},
			ans456{false},
		},

		{
			para456{[]int{3, 1, 4, 2}},
			ans456{true},
		},

		{
			para456{[]int{-1, 3, 2, 0}},
			ans456{true},
		},

		{
			para456{[]int{3, 5, 0, 3, 4}},
			ans456{true},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans456, q.para456
				(find132pattern(p.one))
			}
		}
	}
}

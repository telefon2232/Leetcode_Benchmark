package leetcode

import (
	"testing"
)

type question376 struct {
	para376
	ans376
}

// para 是参数
// one 代表第一个参数
type para376 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans376 struct {
	one int
}

func Benchmark_Problem376(b *testing.B) {

	qs := []question376{

		{
			para376{[]int{1, 7, 4, 9, 2, 5}},
			ans376{6},
		},

		{
			para376{[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}},
			ans376{7},
		},

		{
			para376{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			ans376{2},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans376, q.para376
				(wiggleMaxLength(p.nums))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question453 struct {
	para453
	ans453
}

// para 是参数
// one 代表第一个参数
type para453 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans453 struct {
	one int
}

func Benchmark_Problem453(b *testing.B) {

	qs := []question453{

		{
			para453{[]int{4, 3, 2, 7, 8, 2, 3, 1}},
			ans453{22},
		},

		{
			para453{[]int{1, 2, 3}},
			ans453{3},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans453, q.para453
				(minMoves(p.one))
			}
		}
	}
}

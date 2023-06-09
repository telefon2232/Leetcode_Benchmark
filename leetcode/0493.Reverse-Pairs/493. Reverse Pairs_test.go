package leetcode

import (
	"testing"
)

type question493 struct {
	para493
	ans493
}

// para 是参数
// one 代表第一个参数
type para493 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans493 struct {
	one int
}

func Benchmark_Problem493(b *testing.B) {

	qs := []question493{

		{
			para493{[]int{1, 3, 2, 3, 1}},
			ans493{2},
		},

		{
			para493{[]int{9, 8, 7, 4, 7, 2, 3, 8, 7, 0}},
			ans493{18},
		},

		{
			para493{[]int{2, 4, 3, 5, 1}},
			ans493{3},
		},

		{
			para493{[]int{-5, -5}},
			ans493{1},
		},

		{
			para493{[]int{2147483647, 2147483647, -2147483647, -2147483647, -2147483647, 2147483647}},
			ans493{9},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans493, q.para493
				(reversePairs(p.nums))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question322 struct {
	para322
	ans322
}

// para 是参数
// one 代表第一个参数
type para322 struct {
	one    []int
	amount int
}

// ans 是答案
// one 代表第一个答案
type ans322 struct {
	one int
}

func Benchmark_Problem322(b *testing.B) {

	qs := []question322{

		{
			para322{[]int{186, 419, 83, 408}, 6249},
			ans322{20},
		},

		{
			para322{[]int{1, 2147483647}, 2},
			ans322{2},
		},

		{
			para322{[]int{1, 2, 5}, 11},
			ans322{3},
		},
		{
			para322{[]int{2}, 3},
			ans322{-1},
		},

		{
			para322{[]int{1}, 0},
			ans322{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans322, q.para322
				(coinChange(p.one, p.amount))
			}
		}
	}
}

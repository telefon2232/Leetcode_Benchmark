package leetcode

import (
	"testing"
)

type question605 struct {
	para605
	ans605
}

// para 是参数
// one 代表第一个参数
type para605 struct {
	flowerbed []int
	n         int
}

// ans 是答案
// one 代表第一个答案
type ans605 struct {
	one bool
}

func Benchmark_Problem605(b *testing.B) {

	qs := []question605{

		{
			para605{[]int{1, 0, 0, 0, 1}, 1},
			ans605{true},
		},

		{
			para605{[]int{1, 0, 0, 0, 1}, 2},
			ans605{false},
		},

		{
			para605{[]int{1, 0, 0, 0, 0, 1}, 2},
			ans605{false},
		},

		{
			para605{[]int{0, 0, 1, 0}, 1},
			ans605{true},
		},

		{
			para605{[]int{0, 0, 1, 0, 0}, 1},
			ans605{true},
		},

		{
			para605{[]int{1, 0, 0, 1, 0}, 2},
			ans605{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans605, q.para605
				(canPlaceFlowers(p.flowerbed, p.n))
			}
		}
	}
}

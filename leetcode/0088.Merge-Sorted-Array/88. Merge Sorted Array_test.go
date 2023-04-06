package leetcode

import (
	"testing"
)

type question88 struct {
	para88
	ans88
}

// para 是参数
// one 代表第一个参数
type para88 struct {
	one []int
	m   int
	two []int
	n   int
}

// ans 是答案
// one 代表第一个答案
type ans88 struct {
	one []int
}

func Benchmark_Problem88(b *testing.B) {

	qs := []question88{

		{
			para88{[]int{0}, 0, []int{1}, 1},
			ans88{[]int{1}},
		},

		{
			para88{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
			ans88{[]int{1, 2, 2, 3, 5, 6}},
		},
	}
	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans88, q.para88

				merge(p.one, p.m, p.two, p.n)

			}
		}
	}
}

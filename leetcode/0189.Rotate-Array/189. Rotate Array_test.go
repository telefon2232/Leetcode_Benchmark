package leetcode

import (
	"testing"
)

type question189 struct {
	para189
	ans189
}

// para 是参数
// one 代表第一个参数
type para189 struct {
	nums []int
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans189 struct {
	one []int
}

func Benchmark_Problem189(b *testing.B) {

	qs := []question189{

		{
			para189{[]int{1, 2, 3, 4, 5, 6, 7}, 3},
			ans189{[]int{5, 6, 7, 1, 2, 3, 4}},
		},

		{
			para189{[]int{-1, -100, 3, 99}, 2},
			ans189{[]int{3, 99, -1, -100}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans189, q.para189

				rotate(p.nums, p.k)

			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question747 struct {
	para747
	ans747
}

// para 是参数
// one 代表第一个参数
type para747 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans747 struct {
	one int
}

func Benchmark_Problem747(b *testing.B) {

	qs := []question747{

		{
			para747{[]int{3, 6, 1, 0}},
			ans747{1},
		},

		{
			para747{[]int{1, 2, 3, 4}},
			ans747{-1},
		},

		{
			para747{[]int{1}},
			ans747{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans747, q.para747
				(dominantIndex(p.nums))
			}
		}
	}
}

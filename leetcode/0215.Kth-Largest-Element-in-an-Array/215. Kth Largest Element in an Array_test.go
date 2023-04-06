package leetcode

import (
	"testing"
)

type question215 struct {
	para215
	ans215
}

// para 是参数
// one 代表第一个参数
type para215 struct {
	one []int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans215 struct {
	one int
}

func Benchmark_Problem215(b *testing.B) {

	qs := []question215{

		{
			para215{[]int{3, 2, 1}, 2},
			ans215{2},
		},

		{
			para215{[]int{3, 2, 1, 5, 6, 4}, 2},
			ans215{5},
		},

		{
			para215{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4},
			ans215{4},
		},

		{
			para215{[]int{0, 0, 0, 0, 0}, 2},
			ans215{0},
		},

		{
			para215{[]int{1}, 1},
			ans215{1},
		},

		{
			para215{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}, 20},
			ans215{2},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans215, q.para215
				(findKthLargest(p.one, p.two))
			}
		}
	}
}

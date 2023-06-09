package leetcode

import (
	"testing"
)

type question532 struct {
	para532
	ans532
}

// para 是参数
// one 代表第一个参数
type para532 struct {
	one []int
	n   int
}

// ans 是答案
// one 代表第一个答案
type ans532 struct {
	one int
}

func Benchmark_Problem532(b *testing.B) {

	qs := []question532{

		{
			para532{[]int{3, 1, 4, 1, 5}, 2},
			ans532{2},
		},

		{
			para532{[]int{1, 2, 3, 4, 5}, 1},
			ans532{4},
		},

		{
			para532{[]int{1, 3, 1, 5, 4}, 0},
			ans532{1},
		},

		{
			para532{[]int{}, 3},
			ans532{0},
		},

		// question532{
		// 	para532{[]int{1, 2, 3, 2, 3, 2, 3, 2}, 0},
		// 	ans532{[]int{1, 2, 3, 2, 3, 2, 3, 2}},
		// },

		// question532{
		// 	para532{[]int{1, 2, 3, 4, 5}, 5},
		// 	ans532{[]int{1, 2, 3, 4}},
		// },

		// question532{
		// 	para532{[]int{}, 5},
		// 	ans532{[]int{}},
		// },

		// question532{
		// 	para532{[]int{1, 2, 3, 4, 5}, 10},
		// 	ans532{[]int{1, 2, 3, 4, 5}},
		// },

		// question532{
		// 	para532{[]int{1}, 1},
		// 	ans532{[]int{}},
		// },
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans532, q.para532
				(findPairs(p.one, p.n))
			}
		}
	}
}

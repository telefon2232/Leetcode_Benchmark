package leetcode

import (
	"testing"
)

type question628 struct {
	para628
	ans628
}

// para 是参数
// one 代表第一个参数
type para628 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans628 struct {
	one int
}

func Benchmark_Problem628(b *testing.B) {

	qs := []question628{

		{
			para628{[]int{3, -1, 4}},
			ans628{-12},
		},

		{
			para628{[]int{1, 2}},
			ans628{2},
		},

		{
			para628{[]int{1, 2, 3}},
			ans628{6},
		},

		{
			para628{[]int{1, 2, 3, 4}},
			ans628{24},
		},

		{
			para628{[]int{-2}},
			ans628{-2},
		},

		{
			para628{[]int{0}},
			ans628{0},
		},

		{
			para628{[]int{2, 3, -2, 4}},
			ans628{-24},
		},

		{
			para628{[]int{-2, 0, -1}},
			ans628{0},
		},

		{
			para628{[]int{-2, 0, -1, 2, 3, 1, 10}},
			ans628{60},
		},
	}

	for _, q := range qs {
		_, p := q.ans628, q.para628
		(maximumProduct(p.one))
	}
}

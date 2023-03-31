package leetcode

import (
	"testing"
)

type question1011 struct {
	para1011
	ans1011
}

// para 是参数
// one 代表第一个参数
type para1011 struct {
	weights []int
	D       int
}

// ans 是答案
// one 代表第一个答案
type ans1011 struct {
	one int
}

func Benchmark_Problem1011(b *testing.B) {

	qs := []question1011{

		{
			para1011{[]int{7, 2, 5, 10, 8}, 2},
			ans1011{18},
		},

		{
			para1011{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
			ans1011{15},
		},

		{
			para1011{[]int{3, 2, 2, 4, 1, 4}, 3},
			ans1011{6},
		},

		{
			para1011{[]int{1, 2, 3, 1, 1}, 4},
			ans1011{3},
		},
	}

	for _, q := range qs {
		_, p := q.ans1011, q.para1011
		(shipWithinDays(p.weights, p.D))
	}
}

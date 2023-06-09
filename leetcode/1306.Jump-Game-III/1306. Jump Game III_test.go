package leetcode

import (
	"testing"
)

type question1306 struct {
	para1306
	ans1306
}

// para 是参数
// one 代表第一个参数
type para1306 struct {
	arr   []int
	start int
}

// ans 是答案
// one 代表第一个答案
type ans1306 struct {
	one bool
}

func Benchmark_Problem1306(b *testing.B) {

	qs := []question1306{

		{
			para1306{[]int{4, 2, 3, 0, 3, 1, 2}, 5},
			ans1306{true},
		},

		{
			para1306{[]int{4, 2, 3, 0, 3, 1, 2}, 0},
			ans1306{true},
		},

		{
			para1306{[]int{3, 0, 2, 1, 2}, 2},
			ans1306{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1306, q.para1306
				(canReach(p.arr, p.start))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question1235 struct {
	para1235
	ans1235
}

// para 是参数
// one 代表第一个参数
type para1235 struct {
	startTime []int
	endTime   []int
	profit    []int
}

// ans 是答案
// one 代表第一个答案
type ans1235 struct {
	one int
}

func Benchmark_Problem1235(b *testing.B) {

	qs := []question1235{

		{
			para1235{[]int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70}},
			ans1235{120},
		},

		{
			para1235{[]int{1, 2, 3, 4, 6}, []int{3, 5, 10, 6, 9}, []int{20, 20, 100, 70, 60}},
			ans1235{150},
		},

		{
			para1235{[]int{1, 1, 1}, []int{2, 3, 4}, []int{5, 6, 4}},
			ans1235{6},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1235, q.para1235
				(jobScheduling(p.startTime, p.endTime, p.profit))
			}
		}
	}
}

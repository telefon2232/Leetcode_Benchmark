package leetcode

import (
	"testing"
)

type question1207 struct {
	para1207
	ans1207
}

// para 是参数
// one 代表第一个参数
type para1207 struct {
	arr []int
}

// ans 是答案
// one 代表第一个答案
type ans1207 struct {
	one bool
}

func Benchmark_Problem1207(b *testing.B) {

	qs := []question1207{

		{
			para1207{[]int{1, 2, 2, 1, 1, 3}},
			ans1207{true},
		},

		{
			para1207{[]int{1, 2}},
			ans1207{false},
		},

		{
			para1207{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}},
			ans1207{true},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1207, q.para1207
				(uniqueOccurrences(p.arr))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question1640 struct {
	para1640
	ans1640
}

// para 是参数
// one 代表第一个参数
type para1640 struct {
	arr    []int
	pieces [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1640 struct {
	one bool
}

func Benchmark_Problem1640(b *testing.B) {

	qs := []question1640{

		{
			para1640{[]int{85}, [][]int{{85}}},
			ans1640{true},
		},

		{
			para1640{[]int{15, 88}, [][]int{{88}, {15}}},
			ans1640{true},
		},

		{
			para1640{[]int{49, 18, 16}, [][]int{{16, 18, 49}}},
			ans1640{false},
		},

		{
			para1640{[]int{91, 4, 64, 78}, [][]int{{78}, {4, 64}, {91}}},
			ans1640{true},
		},

		{
			para1640{[]int{1, 3, 5, 7}, [][]int{{2, 4, 6, 8}}},
			ans1640{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1640, q.para1640
				(canFormArray(p.arr, p.pieces))
			}
		}
	}
}

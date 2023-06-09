package leetcode

import (
	"testing"
)

type question1439 struct {
	para1439
	ans1439
}

// para 是参数
// one 代表第一个参数
type para1439 struct {
	mat [][]int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans1439 struct {
	one int
}

func Benchmark_Problem1439(b *testing.B) {

	qs := []question1439{

		{
			para1439{[][]int{{1, 3, 11}, {2, 4, 6}}, 5},
			ans1439{7},
		},

		{
			para1439{[][]int{{1, 3, 11}, {2, 4, 6}}, 9},
			ans1439{17},
		},
		{
			para1439{[][]int{{1, 10, 10}, {1, 4, 5}, {2, 3, 6}}, 7},
			ans1439{9},
		},
		{
			para1439{[][]int{{1, 1, 10}, {2, 2, 9}}, 7},
			ans1439{12},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1439, q.para1439
				(kthSmallest(p.mat, p.k))
			}
		}
	}
}

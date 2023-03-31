package leetcode

import (
	"testing"
)

type question1738 struct {
	para1738
	ans1738
}

// para 是参数
// one 代表第一个参数
type para1738 struct {
	matrix [][]int
	k      int
}

// ans 是答案
// one 代表第一个答案
type ans1738 struct {
	one int
}

func Benchmark_Problem1738(b *testing.B) {

	qs := []question1738{

		{
			para1738{[][]int{{5, 2}, {1, 6}}, 1},
			ans1738{7},
		},

		{
			para1738{[][]int{{5, 2}, {1, 6}}, 2},
			ans1738{5},
		},

		{
			para1738{[][]int{{5, 2}, {1, 6}}, 3},
			ans1738{4},
		},

		{
			para1738{[][]int{{5, 2}, {1, 6}}, 4},
			ans1738{0},
		},
	}

	for _, q := range qs {
		_, p := q.ans1738, q.para1738
		(kthLargestValue(p.matrix, p.k))
	}
}

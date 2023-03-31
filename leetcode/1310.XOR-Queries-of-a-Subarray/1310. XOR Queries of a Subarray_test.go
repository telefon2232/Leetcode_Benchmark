package leetcode

import (
	"testing"
)

type question1310 struct {
	para1310
	ans1310
}

// para 是参数
// one 代表第一个参数
type para1310 struct {
	arr     []int
	queries [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1310 struct {
	one []int
}

func Benchmark_Problem1310(b *testing.B) {

	qs := []question1310{

		{
			para1310{[]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}},
			ans1310{[]int{2, 7, 14, 8}},
		},

		{
			para1310{[]int{4, 8, 2, 10}, [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}},
			ans1310{[]int{8, 0, 4, 4}},
		},
	}

	for _, q := range qs {
		_, p := q.ans1310, q.para1310
		(xorQueries(p.arr, p.queries))
	}
}

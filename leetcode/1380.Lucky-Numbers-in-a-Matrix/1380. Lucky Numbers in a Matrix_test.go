package leetcode

import (
	"testing"
)

type question1380 struct {
	para1380
	ans1380
}

// para 是参数
// one 代表第一个参数
type para1380 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1380 struct {
	one []int
}

func Benchmark_Problem1380(b *testing.B) {

	qs := []question1380{

		{
			para1380{[][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}},
			ans1380{[]int{15}},
		},

		{
			para1380{[][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}},
			ans1380{[]int{12}},
		},

		{
			para1380{[][]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}},
			ans1380{[]int{1}},
		},

		{
			para1380{[][]int{{7, 8}, {1, 2}}},
			ans1380{[]int{7}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1380, q.para1380
				(luckyNumbers(p.one))
			}
		}
	}
}

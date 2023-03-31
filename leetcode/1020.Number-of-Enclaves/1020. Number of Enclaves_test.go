package leetcode

import (
	"testing"
)

type question1020 struct {
	para1020
	ans1020
}

// para 是参数
// one 代表第一个参数
type para1020 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1020 struct {
	one int
}

func Benchmark_Problem1020(b *testing.B) {

	qs := []question1020{

		{
			para1020{[][]int{
				{1, 1, 1, 1, 0},
				{1, 1, 0, 1, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0},
			}},
			ans1020{0},
		},

		{
			para1020{[][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 0, 1, 1},
			}},
			ans1020{1},
		},

		{
			para1020{[][]int{
				{1, 1, 1, 1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0, 1, 1, 0},
				{1, 0, 1, 0, 1, 1, 1, 0},
				{1, 0, 0, 0, 0, 1, 0, 1},
				{1, 1, 1, 1, 1, 1, 1, 0},
			}},
			ans1020{1},
		},

		{
			para1020{[][]int{
				{0, 0, 1, 0, 0},
				{0, 1, 0, 1, 0},
				{0, 1, 1, 1, 0},
			}},
			ans1020{0},
		},

		{
			para1020{[][]int{
				{1, 1, 1, 1, 1, 1, 1},
				{1, 0, 0, 0, 0, 0, 1},
				{1, 0, 1, 1, 1, 0, 1},
				{1, 0, 1, 0, 1, 0, 1},
				{1, 0, 1, 1, 1, 0, 1},
				{1, 0, 0, 0, 0, 0, 1},
				{1, 1, 1, 1, 1, 1, 1},
			}},
			ans1020{8},
		},
	}

	for _, q := range qs {
		_, p := q.ans1020, q.para1020
		(numEnclaves(p.one))
	}
}

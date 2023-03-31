package leetcode

import (
	"testing"
)

type question491 struct {
	para491
	ans491
}

// para 是参数
// one 代表第一个参数
type para491 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans491 struct {
	one [][]int
}

func Benchmark_Problem491(b *testing.B) {

	qs := []question491{

		{
			para491{[]int{4, 3, 2, 1}},
			ans491{[][]int{}},
		},

		{
			para491{[]int{4, 6, 7, 7}},
			ans491{[][]int{{4, 6}, {4, 7}, {4, 6, 7}, {4, 6, 7, 7}, {6, 7}, {6, 7, 7}, {7, 7}, {4, 7, 7}}},
		},
	}

	for _, q := range qs {
		_, p := q.ans491, q.para491
		(findSubsequences(p.one))
	}
}

package leetcode

import (
	"testing"
)

type question35 struct {
	para35
	ans35
}

// para 是参数
// one 代表第一个参数
type para35 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans35 struct {
	one int
}

func Benchmark_Problem35(b *testing.B) {

	qs := []question35{

		{
			para35{[]int{1, 3, 5, 6}, 5},
			ans35{2},
		},

		{
			para35{[]int{1, 3, 5, 6}, 2},
			ans35{1},
		},

		{
			para35{[]int{1, 3, 5, 6}, 7},
			ans35{4},
		},

		{
			para35{[]int{1, 3, 5, 6}, 0},
			ans35{0},
		},
	}

	for _, q := range qs {
		_, p := q.ans35, q.para35
		(searchInsert(p.nums, p.target))
	}
}

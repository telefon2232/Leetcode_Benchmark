package leetcode

import (
	"testing"
)

type question704 struct {
	para704
	ans704
}

// para 是参数
// one 代表第一个参数
type para704 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans704 struct {
	one int
}

func Benchmark_Problem704(b *testing.B) {

	qs := []question704{

		{
			para704{[]int{-1, 0, 3, 5, 9, 12}, 9},
			ans704{4},
		},

		{
			para704{[]int{-1, 0, 3, 5, 9, 12}, 2},
			ans704{-1},
		},
	}

	for _, q := range qs {
		_, p := q.ans704, q.para704
		(search704(p.nums, p.target))
	}
}

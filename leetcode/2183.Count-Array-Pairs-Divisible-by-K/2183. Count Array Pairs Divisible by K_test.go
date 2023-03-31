package leetcode

import (
	"testing"
)

type question2182 struct {
	para2182
	ans2182
}

// para 是参数
// one 代表第一个参数
type para2182 struct {
	nums []int
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans2182 struct {
	one int64
}

func Benchmark_Problem2182(b *testing.B) {

	qs := []question2182{

		{
			para2182{[]int{1, 2, 3, 4, 5}, 2},
			ans2182{7},
		},

		{
			para2182{[]int{1, 2, 3, 4}, 5},
			ans2182{0},
		},
	}

	for _, q := range qs {
		_, p := q.ans2182, q.para2182
		(countPairs(p.nums, p.k))
	}
}

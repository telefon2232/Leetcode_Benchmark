package leetcode

import (
	"testing"
)

type question523 struct {
	para523
	ans523
}

// para 是参数
// one 代表第一个参数
type para523 struct {
	nums []int
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans523 struct {
	one bool
}

func Benchmark_Problem523(b *testing.B) {

	qs := []question523{

		{
			para523{[]int{23, 2, 4, 6, 7}, 6},
			ans523{true},
		},

		{
			para523{[]int{23, 2, 6, 4, 7}, 6},
			ans523{true},
		},

		{
			para523{[]int{23, 2, 6, 4, 7}, 13},
			ans523{false},
		},
	}

	for _, q := range qs {
		_, p := q.ans523, q.para523
		(checkSubarraySum(p.nums, p.k))
	}
}

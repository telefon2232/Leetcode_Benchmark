package leetcode

import (
	"testing"
)

type question228 struct {
	para228
	ans228
}

// para 是参数
// one 代表第一个参数
type para228 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans228 struct {
	ans []string
}

func Benchmark_Problem228(b *testing.B) {

	qs := []question228{

		{
			para228{[]int{0, 1, 2, 4, 5, 7}},
			ans228{[]string{"0->2", "4->5", "7"}},
		},

		{
			para228{[]int{0, 2, 3, 4, 6, 8, 9}},
			ans228{[]string{"0", "2->4", "6", "8->9"}},
		},

		{
			para228{[]int{}},
			ans228{[]string{}},
		},

		{
			para228{[]int{-1}},
			ans228{[]string{"-1"}},
		},

		{
			para228{[]int{0}},
			ans228{[]string{"0"}},
		},
	}

	for _, q := range qs {
		_, p := q.ans228, q.para228
		(summaryRanges(p.nums))
	}
}

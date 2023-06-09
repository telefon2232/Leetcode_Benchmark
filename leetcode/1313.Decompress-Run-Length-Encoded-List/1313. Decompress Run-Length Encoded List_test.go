package leetcode

import (
	"testing"
)

type question1313 struct {
	para1313
	ans1313
}

// para 是参数
// one 代表第一个参数
type para1313 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans1313 struct {
	one []int
}

func Benchmark_Problem1313(b *testing.B) {

	qs := []question1313{

		{
			para1313{[]int{1, 2, 3, 4}},
			ans1313{[]int{2, 4, 4, 4}},
		},

		{
			para1313{[]int{1, 1, 2, 3}},
			ans1313{[]int{1, 3, 3}},
		},

		{
			para1313{[]int{}},
			ans1313{[]int{}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1313, q.para1313

				(decompressRLElist(p.nums))
			}
		}
	}
}

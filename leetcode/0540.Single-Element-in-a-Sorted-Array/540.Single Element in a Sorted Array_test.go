package leetcode

import (
	"testing"
)

type question540 struct {
	para540
	ans540
}

// para 是参数
type para540 struct {
	nums []int
}

// ans 是答案
type ans540 struct {
	ans int
}

func Benchmark_Problem540(b *testing.B) {

	qs := []question540{

		{
			para540{[]int{1, 1, 2, 3, 3, 4, 4, 8, 8}},
			ans540{2},
		},

		{
			para540{[]int{3, 3, 7, 7, 10, 11, 11}},
			ans540{10},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans540, q.para540

				(singleNonDuplicate(p.nums))
			}
		}
	}
}

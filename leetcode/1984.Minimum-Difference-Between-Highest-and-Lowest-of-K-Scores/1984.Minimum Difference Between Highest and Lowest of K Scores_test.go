package leetcode

import (
	"testing"
)

type question1984 struct {
	para1984
	ans1984
}

// para 是参数
type para1984 struct {
	nums []int
	k    int
}

// ans 是答案
type ans1984 struct {
	ans int
}

func Benchmark_Problem1984(b *testing.B) {

	qs := []question1984{

		{
			para1984{[]int{90}, 1},
			ans1984{0},
		},

		{
			para1984{[]int{9, 4, 1, 7}, 2},
			ans1984{2},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1984, q.para1984

				(minimumDifference(p.nums, p.k))
			}
		}
	}
}

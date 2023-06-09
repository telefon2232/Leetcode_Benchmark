package leetcode

import (
	"testing"
)

type question1104 struct {
	para1104
	ans1104
}

// para 是参数
type para1104 struct {
	label int
}

// ans 是答案
type ans1104 struct {
	ans []int
}

func Benchmark_Problem1104(b *testing.B) {

	qs := []question1104{

		{
			para1104{14},
			ans1104{[]int{1, 3, 4, 14}},
		},

		{
			para1104{26},
			ans1104{[]int{1, 2, 6, 10, 26}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1104, q.para1104
				(pathInZigZagTree(p.label))
			}
		}
	}
}

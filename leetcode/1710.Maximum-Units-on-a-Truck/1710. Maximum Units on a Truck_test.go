package leetcode

import (
	"testing"
)

type question1710 struct {
	para1710
	ans1710
}

// para 是参数
// one 代表第一个参数
type para1710 struct {
	boxTypes  [][]int
	truckSize int
}

// ans 是答案
// one 代表第一个答案
type ans1710 struct {
	one int
}

func Benchmark_Problem1710(b *testing.B) {

	qs := []question1710{

		{
			para1710{[][]int{{1, 3}, {2, 2}, {3, 1}}, 4},
			ans1710{8},
		},

		{
			para1710{[][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10},
			ans1710{91},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1710, q.para1710
				(maximumUnits(p.boxTypes, p.truckSize))
			}
		}
	}
}

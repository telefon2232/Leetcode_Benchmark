package leetcode

import (
	"testing"
)

type question1690 struct {
	para1690
	ans1690
}

// para 是参数
// one 代表第一个参数
type para1690 struct {
	stones []int
}

// ans 是答案
// one 代表第一个答案
type ans1690 struct {
	one int
}

func Benchmark_Problem1690(b *testing.B) {

	qs := []question1690{

		{
			para1690{[]int{5, 3, 1, 4, 2}},
			ans1690{6},
		},

		{
			para1690{[]int{7, 90, 5, 1, 100, 10, 10, 2}},
			ans1690{122},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1690, q.para1690
				(stoneGameVII(p.stones))
			}
		}
	}
}

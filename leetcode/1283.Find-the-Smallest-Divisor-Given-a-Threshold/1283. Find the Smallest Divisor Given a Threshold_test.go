package leetcode

import (
	"testing"
)

type question1283 struct {
	para1283
	ans1283
}

// para 是参数
// one 代表第一个参数
type para1283 struct {
	nums      []int
	threshold int
}

// ans 是答案
// one 代表第一个答案
type ans1283 struct {
	one int
}

func Benchmark_Problem1283(b *testing.B) {

	qs := []question1283{

		{
			para1283{[]int{1, 2, 5, 9}, 6},
			ans1283{5},
		},

		{
			para1283{[]int{2, 3, 5, 7, 11}, 11},
			ans1283{3},
		},

		{
			para1283{[]int{19}, 5},
			ans1283{4},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1283, q.para1283
				(smallestDivisor(p.nums, p.threshold))
			}
		}
	}
}

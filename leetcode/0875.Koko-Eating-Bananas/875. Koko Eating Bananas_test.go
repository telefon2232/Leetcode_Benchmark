package leetcode

import (
	"testing"
)

type question875 struct {
	para875
	ans875
}

// para 是参数
// one 代表第一个参数
type para875 struct {
	piles []int
	H     int
}

// ans 是答案
// one 代表第一个答案
type ans875 struct {
	one int
}

func Benchmark_Problem875(b *testing.B) {

	qs := []question875{

		{
			para875{[]int{3, 6, 7, 11}, 8},
			ans875{4},
		},

		{
			para875{[]int{30, 11, 23, 4, 20}, 5},
			ans875{30},
		},

		{
			para875{[]int{30, 11, 23, 4, 20}, 6},
			ans875{23},
		},

		{
			para875{[]int{4, 9, 11, 17}, 8},
			ans875{6},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans875, q.para875
		(minEatingSpeed(p.piles, p.H))
	}
}}}

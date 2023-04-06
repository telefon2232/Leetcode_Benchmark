package leetcode

import (
	"testing"
)

type question910 struct {
	para910
	ans910
}

type para910 struct {
	A []int
	K int
}

type ans910 struct {
	one int
}

// Test_Problem910 ...
func Benchmark_Problem910(b *testing.B) {

	qs := []question910{

		{
			para910{[]int{1}, 0},
			ans910{0},
		},
		{
			para910{[]int{0, 10}, 2},
			ans910{6},
		},
		{
			para910{[]int{1, 3, 6}, 3},
			ans910{3},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans910, q.para910
				(smallestRangeII(p.A, p.K))
			}
		}
	}
}

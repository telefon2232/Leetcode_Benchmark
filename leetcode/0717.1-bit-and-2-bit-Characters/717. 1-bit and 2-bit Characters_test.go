package leetcode

import (
	"testing"
)

type question717 struct {
	para717
	ans717
}

// para 是参数
// one 代表第一个参数
type para717 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans717 struct {
	one bool
}

func Benchmark_Problem717(b *testing.B) {

	qs := []question717{

		{
			para717{[]int{1, 0, 0}},
			ans717{true},
		},

		{
			para717{[]int{1, 1, 1, 0}},
			ans717{false},
		},

		{
			para717{[]int{0, 1, 1, 1, 0, 0}},
			ans717{true},
		},

		{
			para717{[]int{1, 1, 1, 1, 0}},
			ans717{true},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans717, q.para717
				(isOneBitCharacter(p.one))
			}
		}
	}
}

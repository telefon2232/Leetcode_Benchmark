package leetcode

import (
	"testing"
)

type question781 struct {
	para781
	ans781
}

// para 是参数
// one 代表第一个参数
type para781 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans781 struct {
	one int
}

func Benchmark_Problem781(b *testing.B) {

	qs := []question781{
		{
			para781{[]int{1, 1, 2}},
			ans781{5},
		},

		{
			para781{[]int{10, 10, 10}},
			ans781{11},
		},

		{
			para781{[]int{}},
			ans781{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans781, q.para781
				(numRabbits(p.one))
			}
		}
	}
}

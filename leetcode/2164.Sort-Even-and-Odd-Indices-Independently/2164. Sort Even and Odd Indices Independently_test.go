package leetcode

import (
	"testing"
)

type question2164 struct {
	para2164
	ans2164
}

// para 是参数
// one 代表第一个参数
type para2164 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans2164 struct {
	one []int
}

func Benchmark_Problem1(b *testing.B) {

	qs := []question2164{
		{
			para2164{[]int{4, 1, 2, 3}},
			ans2164{[]int{2, 3, 4, 1}},
		},

		{
			para2164{[]int{2, 1}},
			ans2164{[]int{2, 1}},
		},
	}

	for _, q := range qs {
		_, p := q.ans2164, q.para2164
		(sortEvenOdd(p.nums))
	}
}

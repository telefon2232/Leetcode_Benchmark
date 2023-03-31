package leetcode

import (
	"testing"
)

type question826 struct {
	para826
	ans826
}

// para 是参数
// one 代表第一个参数
type para826 struct {
	one   []int
	two   []int
	three []int
}

// ans 是答案
// one 代表第一个答案
type ans826 struct {
	one int
}

func Benchmark_Problem826(b *testing.B) {

	qs := []question826{

		{
			para826{[]int{2, 4, 6, 8, 10}, []int{10, 20, 30, 40, 50}, []int{4, 5, 6, 7}},
			ans826{100},
		},

		{
			para826{[]int{85, 47, 57}, []int{24, 66, 99}, []int{40, 25, 25}},
			ans826{0},
		},
	}

	for _, q := range qs {
		_, p := q.ans826, q.para826
		(maxProfitAssignment(p.one, p.two, p.three))
	}
}

package leetcode

import (
	"testing"
)

type question611 struct {
	para611
	ans611
}

// para 是参数
// one 代表第一个参数
type para611 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans611 struct {
	one int
}

func Benchmark_Problem611(b *testing.B) {

	qs := []question611{

		{
			para611{[]int{2, 2, 3, 4}},
			ans611{3},
		},

		{
			para611{[]int{4, 2, 3, 4}},
			ans611{4},
		},

		{
			para611{[]int{0, 0}},
			ans611{0},
		},
	}

	for _, q := range qs {
		_, p := q.ans611, q.para611
		(triangleNumber(p.nums))
	}
}

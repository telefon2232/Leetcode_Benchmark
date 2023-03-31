package leetcode

import (
	"testing"
)

type question414 struct {
	para414
	ans414
}

// para 是参数
// one 代表第一个参数
type para414 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans414 struct {
	one int
}

func Benchmark_Problem414(b *testing.B) {

	qs := []question414{

		{
			para414{[]int{1, 1, 2}},
			ans414{2},
		},

		{
			para414{[]int{3, 2, 1}},
			ans414{1},
		},

		{
			para414{[]int{1, 2}},
			ans414{2},
		},

		{
			para414{[]int{2, 2, 3, 1}},
			ans414{1},
		},
	}

	for _, q := range qs {
		_, p := q.ans414, q.para414
		(thirdMax(p.one))
	}
}

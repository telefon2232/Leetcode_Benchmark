package leetcode

import (
	"testing"
)

type question594 struct {
	para594
	ans594
}

// para 是参数
// one 代表第一个参数
type para594 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans594 struct {
	one int
}

func Benchmark_Problem594(b *testing.B) {

	qs := []question594{

		{
			para594{[]int{1, 3, 2, 2, 5, 2, 3, 7}},
			ans594{5},
		},
	}

	for _, q := range qs {
		_, p := q.ans594, q.para594
		(findLHS(p.one))
	}
}

package leetcode

import (
	"testing"
)

type question575 struct {
	para575
	ans575
}

// para 是参数
// one 代表第一个参数
type para575 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans575 struct {
	one int
}

func Benchmark_Problem575(b *testing.B) {

	qs := []question575{

		{
			para575{[]int{1, 1, 2, 2, 3, 3}},
			ans575{3},
		},

		{
			para575{[]int{1, 1, 2, 3}},
			ans575{2},
		},
	}

	for _, q := range qs {
		_, p := q.ans575, q.para575
		(distributeCandies(p.one))
	}
}

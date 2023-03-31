package leetcode

import (
	"testing"
)

type question137 struct {
	para137
	ans137
}

// para 是参数
// one 代表第一个参数
type para137 struct {
	s []int
}

// ans 是答案
// one 代表第一个答案
type ans137 struct {
	one int
}

func Benchmark_Problem137(b *testing.B) {

	qs := []question137{

		{
			para137{[]int{2, 2, 3, 2}},
			ans137{3},
		},

		{
			para137{[]int{0, 1, 0, 1, 0, 1, 99}},
			ans137{99},
		},
	}

	for _, q := range qs {
		_, p := q.ans137, q.para137
		(singleNumberII(p.s))
	}
}

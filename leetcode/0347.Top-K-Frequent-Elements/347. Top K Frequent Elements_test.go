package leetcode

import (
	"testing"
)

type question347 struct {
	para347
	ans347
}

// para 是参数
// one 代表第一个参数
type para347 struct {
	one []int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans347 struct {
	one []int
}

func Benchmark_Problem347(b *testing.B) {

	qs := []question347{

		{
			para347{[]int{1, 1, 1, 2, 2, 3}, 2},
			ans347{[]int{1, 2}},
		},

		{
			para347{[]int{1}, 1},
			ans347{[]int{1}},
		},
	}

	for _, q := range qs {
		_, p := q.ans347, q.para347
		(topKFrequent(p.one, p.two))
	}
}

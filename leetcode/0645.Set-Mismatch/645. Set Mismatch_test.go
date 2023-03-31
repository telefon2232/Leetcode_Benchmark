package leetcode

import (
	"testing"
)

type question645 struct {
	para645
	ans645
}

// para 是参数
// one 代表第一个参数
type para645 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans645 struct {
	one []int
}

func Benchmark_Problem645(b *testing.B) {

	qs := []question645{

		{
			para645{[]int{1, 2, 2, 4}},
			ans645{[]int{2, 3}},
		},
	}

	for _, q := range qs {
		_, p := q.ans645, q.para645
		(findErrorNums(p.one))
	}
}

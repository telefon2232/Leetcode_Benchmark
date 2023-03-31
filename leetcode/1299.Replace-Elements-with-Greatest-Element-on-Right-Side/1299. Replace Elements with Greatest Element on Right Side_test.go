package leetcode

import (
	"testing"
)

type question1299 struct {
	para1299
	ans1299
}

// para 是参数
// one 代表第一个参数
type para1299 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans1299 struct {
	one []int
}

func Benchmark_Problem1299(b *testing.B) {

	qs := []question1299{

		{
			para1299{[]int{17, 18, 5, 4, 6, 1}},
			ans1299{[]int{18, 6, 6, 6, 1, -1}},
		},
	}

	for _, q := range qs {
		_, p := q.ans1299, q.para1299
		(replaceElements(p.one))
	}
}

package leetcode

import (
	"testing"
)

type question1295 struct {
	para1295
	ans1295
}

// para 是参数
// one 代表第一个参数
type para1295 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans1295 struct {
	one int
}

func Benchmark_Problem1295(b *testing.B) {

	qs := []question1295{

		{
			para1295{[]int{12, 345, 2, 6, 7896}},
			ans1295{2},
		},

		{
			para1295{[]int{555, 901, 482, 1771}},
			ans1295{1},
		},
	}

	for _, q := range qs {
		_, p := q.ans1295, q.para1295
		(findNumbers(p.one))
	}
}

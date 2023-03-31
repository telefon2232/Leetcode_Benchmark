package leetcode

import (
	"testing"
)

type question1700 struct {
	para1700
	ans1700
}

// para 是参数
// one 代表第一个参数
type para1700 struct {
	students   []int
	sandwiches []int
}

// ans 是答案
// one 代表第一个答案
type ans1700 struct {
	one int
}

func Benchmark_Problem1700(b *testing.B) {

	qs := []question1700{

		{
			para1700{[]int{1, 1, 0, 0}, []int{0, 1, 0, 1}},
			ans1700{0},
		},

		{
			para1700{[]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}},
			ans1700{3},
		},
	}

	for _, q := range qs {
		_, p := q.ans1700, q.para1700
		(countStudents(p.students, p.sandwiches))
	}
}

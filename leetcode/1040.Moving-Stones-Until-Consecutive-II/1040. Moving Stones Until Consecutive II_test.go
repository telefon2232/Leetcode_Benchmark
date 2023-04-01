package leetcode

import (
	"testing"
)

type question1040 struct {
	para1040
	ans1040
}

// para 是参数
// one 代表第一个参数
type para1040 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans1040 struct {
	one []int
}

func Benchmark_Problem1040(b *testing.B) {

	qs := []question1040{
		{
			para1040{[]int{7, 4, 9}},
			ans1040{[]int{1, 2}},
		},

		{
			para1040{[]int{6, 5, 4, 3, 10}},
			ans1040{[]int{2, 3}},
		},

		{
			para1040{[]int{100, 101, 104, 102, 103}},
			ans1040{[]int{0, 0}},
		},

		{
			para1040{[]int{1, 3, 5, 7, 10}},
			ans1040{[]int{2, 4}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans1040, q.para1040
		(numMovesStonesII(p.one))
	}
}}}

package leetcode

import (
	"testing"
)

type question1720 struct {
	para1720
	ans1720
}

// para 是参数
// one 代表第一个参数
type para1720 struct {
	encoded []int
	first   int
}

// ans 是答案
// one 代表第一个答案
type ans1720 struct {
	one []int
}

func Benchmark_Problem1720(b *testing.B) {

	qs := []question1720{

		{
			para1720{[]int{1, 2, 3}, 1},
			ans1720{[]int{1, 0, 2, 1}},
		},

		{
			para1720{[]int{6, 2, 7, 3}, 4},
			ans1720{[]int{4, 2, 0, 7, 4}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans1720, q.para1720
		(decode(p.encoded, p.first))
	}
}}}

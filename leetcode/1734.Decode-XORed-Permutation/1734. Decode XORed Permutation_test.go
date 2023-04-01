package leetcode

import (
	"testing"
)

type question1734 struct {
	para1734
	ans1734
}

// para 是参数
// one 代表第一个参数
type para1734 struct {
	encoded []int
}

// ans 是答案
// one 代表第一个答案
type ans1734 struct {
	one []int
}

func Benchmark_Problem1734(b *testing.B) {

	qs := []question1734{

		{
			para1734{[]int{3, 1}},
			ans1734{[]int{1, 2, 3}},
		},

		{
			para1734{[]int{6, 5, 4, 6}},
			ans1734{[]int{2, 4, 1, 5, 3}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans1734, q.para1734
		(decode(p.encoded))
	}
}}}

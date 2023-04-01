package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question1171 struct {
	para1171
	ans1171
}

// para 是参数
// one 代表第一个参数
type para1171 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans1171 struct {
	one []int
}

func Benchmark_Problem1171(b *testing.B) {

	qs := []question1171{

		{
			para1171{[]int{1, 2, -3, 3, 1}},
			ans1171{[]int{3, 1}},
		},

		{
			para1171{[]int{1, 2, 3, -3, 4}},
			ans1171{[]int{1, 2, 4}},
		},

		{
			para1171{[]int{1, 2, 3, -3, -2}},
			ans1171{[]int{1}},
		},

		{
			para1171{[]int{1, -1}},
			ans1171{[]int{}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans1171, q.para1171
		(structures.List2Ints(removeZeroSumSublists(structures.Ints2List(p.one))))
	}
}}}

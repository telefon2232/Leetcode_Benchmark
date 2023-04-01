package leetcode

import (
	"testing"
)

type question1696 struct {
	para1696
	ans1696
}

// para 是参数
// one 代表第一个参数
type para1696 struct {
	nums []int
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans1696 struct {
	one int
}

func Benchmark_Problem1696(b *testing.B) {

	qs := []question1696{

		{
			para1696{[]int{1, -1, -2, 4, -7, 3}, 2},
			ans1696{7},
		},

		{
			para1696{[]int{10, -5, -2, 4, 0, 3}, 3},
			ans1696{17},
		},

		{
			para1696{[]int{1, -5, -20, 4, -1, 3, -6, -3}, 2},
			ans1696{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans1696, q.para1696
		(maxResult(p.nums, p.k))
	}
}}}

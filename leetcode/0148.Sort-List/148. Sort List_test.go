package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question148 struct {
	para148
	ans148
}

// para 是参数
// one 代表第一个参数
type para148 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans148 struct {
	one []int
}

func Benchmark_Problem148(b *testing.B) {

	qs := []question148{

		{
			para148{[]int{1, 2, 3, 4, 5}},
			ans148{[]int{1, 2, 3, 4, 5}},
		},
		{
			para148{[]int{1, 1, 2, 5, 5, 4, 10, 0}},
			ans148{[]int{0, 1, 1, 2, 4, 5, 5}},
		},

		{
			para148{[]int{1}},
			ans148{[]int{1}},
		},

		{
			para148{[]int{}},
			ans148{[]int{}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans148, q.para148
		(structures.List2Ints(sortList(structures.Ints2List(p.one))))
	}
}}}

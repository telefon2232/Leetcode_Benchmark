package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question147 struct {
	para147
	ans147
}

// para 是参数
// one 代表第一个参数
type para147 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans147 struct {
	one []int
}

func Benchmark_Problem147(b *testing.B) {

	qs := []question147{

		{
			para147{[]int{4, 2, 1, 3}},
			ans147{[]int{1, 2, 3, 4}},
		},
		{
			para147{[]int{1}},
			ans147{[]int{1}},
		},

		{
			para147{[]int{-1, 5, 3, 4, 0}},
			ans147{[]int{-1, 0, 3, 4, 5}},
		},

		{
			para147{[]int{}},
			ans147{[]int{}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans147, q.para147
		(structures.List2Ints(insertionSortList(structures.Ints2List(p.one))))
	}
}}}

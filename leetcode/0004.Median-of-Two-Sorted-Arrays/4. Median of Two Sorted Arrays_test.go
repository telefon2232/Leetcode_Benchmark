package leetcode

import (
	"testing"
)

type question4 struct {
	para4
	ans4
}

// para 是参数
// one 代表第一个参数
type para4 struct {
	nums1 []int
	nums2 []int
}

// ans 是答案
// one 代表第一个答案
type ans4 struct {
	one float64
}

func Benchmark_Problem4(b *testing.B) {

	qs := []question4{

		{
			para4{[]int{1, 3}, []int{2}},
			ans4{2.0},
		},

		{
			para4{[]int{1, 2}, []int{3, 4}},
			ans4{2.5},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans4, q.para4
		(findMedianSortedArrays(p.nums1, p.nums2))
	}
}}}

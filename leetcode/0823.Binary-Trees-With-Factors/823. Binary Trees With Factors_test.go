package leetcode

import (
	"testing"
)

type question823 struct {
	para823
	ans823
}

// para 是参数
// one 代表第一个参数
type para823 struct {
	arr []int
}

// ans 是答案
// one 代表第一个答案
type ans823 struct {
	one int
}

func Benchmark_Problem823(b *testing.B) {

	qs := []question823{

		{
			para823{[]int{2, 4}},
			ans823{3},
		},

		{
			para823{[]int{2, 4, 5, 10}},
			ans823{7},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans823, q.para823
				(numFactoredBinaryTrees(p.arr))
			}
		}
	}
}

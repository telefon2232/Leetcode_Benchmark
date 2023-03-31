package leetcode

import (
	"testing"
)

type question64 struct {
	para64
	ans64
}

// para 是参数
// one 代表第一个参数
type para64 struct {
	og [][]int
}

// ans 是答案
// one 代表第一个答案
type ans64 struct {
	one int
}

func Benchmark_Problem64(b *testing.B) {

	qs := []question64{

		{
			para64{[][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			}},
			ans64{7},
		},
	}

	for _, q := range qs {
		_, p := q.ans64, q.para64
		(minPathSum(p.og))
	}
}

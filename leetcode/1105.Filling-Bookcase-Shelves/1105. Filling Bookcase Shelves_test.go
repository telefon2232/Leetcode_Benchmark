package leetcode

import (
	"testing"
)

type question1105 struct {
	para1105
	ans1105
}

// para 是参数
// one 代表第一个参数
type para1105 struct {
	one [][]int
	w   int
}

// ans 是答案
// one 代表第一个答案
type ans1105 struct {
	one int
}

func Benchmark_Problem1105(b *testing.B) {

	qs := []question1105{

		{
			para1105{[][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}, 4},
			ans1105{6},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1105, q.para1105
				(minHeightShelves(p.one, p.w))
			}
		}
	}
}

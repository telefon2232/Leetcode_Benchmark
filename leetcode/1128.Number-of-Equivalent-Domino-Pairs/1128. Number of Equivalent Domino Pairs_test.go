package leetcode

import (
	"testing"
)

type question1128 struct {
	para1128
	ans1128
}

// para 是参数
// one 代表第一个参数
type para1128 struct {
	dominoes [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1128 struct {
	one int
}

func Benchmark_Problem1128(b *testing.B) {

	qs := []question1128{

		{
			para1128{[][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}},
			ans1128{1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1128, q.para1128
				(numEquivDominoPairs(p.dominoes))
			}
		}
	}
}

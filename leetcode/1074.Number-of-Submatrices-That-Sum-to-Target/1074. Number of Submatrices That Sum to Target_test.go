package leetcode

import (
	"testing"
)

type question1074 struct {
	para1074
	ans1074
}

// para 是参数
// one 代表第一个参数
type para1074 struct {
	one [][]int
	t   int
}

// ans 是答案
// one 代表第一个答案
type ans1074 struct {
	one int
}

func Benchmark_Problem1074(b *testing.B) {

	qs := []question1074{

		{
			para1074{[][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}, 0},
			ans1074{4},
		},

		{
			para1074{[][]int{{1, -1}, {-1, 1}}, 0},
			ans1074{5},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1074, q.para1074
				(numSubmatrixSumTarget1(p.one, p.t))
			}
		}
	}
}

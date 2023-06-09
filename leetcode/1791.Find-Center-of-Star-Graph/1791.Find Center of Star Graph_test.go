package leetcode

import (
	"testing"
)

type question1791 struct {
	para1791
	ans1791
}

// para 是参数
type para1791 struct {
	edges [][]int
}

// ans 是答案
type ans1791 struct {
	ans int
}

func Benchmark_Problem1791(b *testing.B) {

	qs := []question1791{

		{
			para1791{[][]int{{1, 2}, {2, 3}, {4, 2}}},
			ans1791{2},
		},

		{
			para1791{[][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}},
			ans1791{1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1791, q.para1791

				(findCenter(p.edges))
			}
		}
	}
}

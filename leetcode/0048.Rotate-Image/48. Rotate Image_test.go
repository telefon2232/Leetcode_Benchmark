package leetcode

import (
	"testing"
)

type question48 struct {
	para48
	ans48
}

// para 是参数
// one 代表第一个参数
type para48 struct {
	s [][]int
}

// ans 是答案
// one 代表第一个答案
type ans48 struct {
	s [][]int
}

func Benchmark_Problem48(b *testing.B) {

	qs := []question48{

		{
			para48{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			ans48{[][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		},

		{
			para48{[][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}},
			ans48{[][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans48, q.para48

				rotate(p.s)

			}
		}
	}
}

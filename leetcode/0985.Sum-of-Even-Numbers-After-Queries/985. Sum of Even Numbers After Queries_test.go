package leetcode

import (
	"testing"
)

type question985 struct {
	para985
	ans985
}

// para 是参数
// one 代表第一个参数
type para985 struct {
	A       []int
	queries [][]int
}

// ans 是答案
// one 代表第一个答案
type ans985 struct {
	one []int
}

func Benchmark_Problem985(b *testing.B) {

	qs := []question985{

		{
			para985{[]int{1, 2, 3, 4}, [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}}},
			ans985{[]int{8, 6, 2, 4}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans985, q.para985
				(sumEvenAfterQueries(p.A, p.queries))
			}
		}
	}
}

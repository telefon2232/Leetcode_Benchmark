package leetcode

import (
	"testing"
)

type question120 struct {
	para120
	ans120
}

// para 是参数
// one 代表第一个参数
type para120 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans120 struct {
	one int
}

func Benchmark_Problem120(b *testing.B) {

	qs := []question120{
		{
			para120{[][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}},
			ans120{11},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans120, q.para120
				(minimumTotal(p.one))
			}
		}
	}
}

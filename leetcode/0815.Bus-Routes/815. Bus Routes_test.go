package leetcode

import (
	"testing"
)

type question815 struct {
	para815
	ans815
}

// para 是参数
// one 代表第一个参数
type para815 struct {
	r [][]int
	s int
	t int
}

// ans 是答案
// one 代表第一个答案
type ans815 struct {
	one int
}

func Benchmark_Problem815(b *testing.B) {

	qs := []question815{

		{
			para815{[][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6},
			ans815{2},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans815, q.para815
				(numBusesToDestination(p.r, p.s, p.t))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question881 struct {
	para881
	ans881
}

// para 是参数
// one 代表第一个参数
type para881 struct {
	s []int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans881 struct {
	one int
}

func Benchmark_Problem881(b *testing.B) {

	qs := []question881{

		{
			para881{[]int{1, 2}, 3},
			ans881{1},
		},

		{
			para881{[]int{3, 2, 2, 1}, 3},
			ans881{3},
		},

		{
			para881{[]int{3, 5, 3, 4}, 5},
			ans881{4},
		},

		{
			para881{[]int{5, 1, 4, 2}, 6},
			ans881{2},
		},

		{
			para881{[]int{3, 2, 2, 1}, 3},
			ans881{3},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans881, q.para881
				(numRescueBoats(p.s, p.k))
			}
		}
	}
}

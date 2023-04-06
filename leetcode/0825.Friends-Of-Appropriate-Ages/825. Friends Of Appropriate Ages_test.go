package leetcocde

import (
	"testing"
)

type question825 struct {
	para825
	ans825
}

// para 是参数
// one 代表第一个参数
type para825 struct {
	ages []int
}

// ans 是答案
// one 代表第一个答案
type ans825 struct {
	one int
}

func Benchmark_Problem825(b *testing.B) {

	qs := []question825{

		{
			para825{[]int{16, 16}},
			ans825{2},
		},

		{
			para825{[]int{16, 17, 18}},
			ans825{2},
		},

		{
			para825{[]int{20, 30, 100, 110, 120}},
			ans825{3},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans825, q.para825
				(numFriendRequests(p.ages))
			}
		}
	}
}

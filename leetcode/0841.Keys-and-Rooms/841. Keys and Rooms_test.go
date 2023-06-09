package leetcode

import (
	"testing"
)

type question841 struct {
	para841
	ans841
}

// para 是参数
// one 代表第一个参数
type para841 struct {
	rooms [][]int
}

// ans 是答案
// one 代表第一个答案
type ans841 struct {
	one bool
}

func Benchmark_Problem841(b *testing.B) {

	qs := []question841{

		{
			para841{[][]int{{1}, {2}, {3}, {}}},
			ans841{true},
		},

		{
			para841{[][]int{{1, 3}, {3, 0, 1}, {2}, {0}}},
			ans841{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans841, q.para841
				(canVisitAllRooms(p.rooms))
			}
		}
	}
}

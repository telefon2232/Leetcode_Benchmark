package leetcode

import (
	"testing"
)

type question2037 struct {
	para2037
	ans2037
}

// para 是参数
type para2037 struct {
	seats    []int
	students []int
}

// ans 是答案
type ans2037 struct {
	ans int
}

func Benchmark_Problem2037(b *testing.B) {

	qs := []question2037{

		{
			para2037{[]int{3, 1, 5}, []int{2, 7, 4}},
			ans2037{4},
		},

		{
			para2037{[]int{4, 1, 5, 9}, []int{1, 3, 2, 6}},
			ans2037{7},
		},

		{
			para2037{[]int{2, 2, 6, 6}, []int{1, 3, 2, 6}},
			ans2037{4},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans2037, q.para2037

				(minMovesToSeat(p.seats, p.students))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question219 struct {
	para219
	ans219
}

// para 是参数
// one 代表第一个参数
type para219 struct {
	one []int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans219 struct {
	one bool
}

func Benchmark_Problem219(b *testing.B) {

	qs := []question219{

		{
			para219{[]int{1, 2, 3, 1}, 3},
			ans219{true},
		},

		{
			para219{[]int{1, 0, 0, 1}, 1},
			ans219{true},
		},

		{
			para219{[]int{1, 2, 3, 1, 2, 3}, 2},
			ans219{false},
		},
	}

	for _, q := range qs {
		_, p := q.ans219, q.para219
		(containsNearbyDuplicate(p.one, p.k))
	}
}

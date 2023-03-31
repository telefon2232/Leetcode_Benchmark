package leetcode

import (
	"testing"
)

type question27 struct {
	para27
	ans27
}

// para 是参数
// one 代表第一个参数
type para27 struct {
	one []int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans27 struct {
	one int
}

func Benchmark_Problem27(b *testing.B) {

	qs := []question27{

		{
			para27{[]int{1, 0, 1}, 1},
			ans27{1},
		},

		{
			para27{[]int{0, 1, 0, 3, 0, 12}, 0},
			ans27{3},
		},

		{
			para27{[]int{0, 1, 0, 3, 0, 0, 0, 0, 1, 12}, 0},
			ans27{4},
		},

		{
			para27{[]int{0, 0, 0, 0, 0}, 0},
			ans27{0},
		},

		{
			para27{[]int{1}, 1},
			ans27{0},
		},

		{
			para27{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2},
			ans27{5},
		},
	}

	for _, q := range qs {
		_, p := q.ans27, q.para27
		(removeElement(p.one, p.two))
	}
}

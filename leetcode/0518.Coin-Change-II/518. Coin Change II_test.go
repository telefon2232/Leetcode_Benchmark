package leetcode

import (
	"testing"
)

type question518 struct {
	para518
	ans518
}

// para 是参数
// one 代表第一个参数
type para518 struct {
	amount int
	coins  []int
}

// ans 是答案
// one 代表第一个答案
type ans518 struct {
	one int
}

func Benchmark_Problem518(b *testing.B) {

	qs := []question518{

		{
			para518{5, []int{1, 2, 5}},
			ans518{4},
		},

		{
			para518{3, []int{2}},
			ans518{0},
		},

		{
			para518{10, []int{10}},
			ans518{1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans518, q.para518
				(change(p.amount, p.coins))
			}
		}
	}
}

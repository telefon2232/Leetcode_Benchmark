package leetcode

import (
	"testing"
)

type question710 struct {
	para710
	ans710
}

// para 是参数
// one 代表第一个参数
type para710 struct {
	n         int
	blackList []int
	times     int
}

// ans 是答案
// one 代表第一个答案
type ans710 struct {
	one []int
}

func Benchmark_Problem710(b *testing.B) {

	qs := []question710{

		{
			para710{1, []int{}, 3},
			ans710{[]int{0, 0, 0}},
		},

		{
			para710{2, []int{}, 3},
			ans710{[]int{1, 1, 1}},
		},

		{
			para710{3, []int{1}, 3},
			ans710{[]int{0, 0, 2}},
		},

		{
			para710{4, []int{2}, 3},
			ans710{[]int{1, 3, 1}},
		},

		{
			para710{10000000, []int{1, 9999, 999999, 99999, 100, 0}, 10},
			ans710{[]int{400, 200, 300}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans710, q.para710

				obj := Constructor710(p.n, p.blackList)

				for i := 0; i < p.times; i++ {
					(obj.Pick())
				}

			}
		}
	}
}

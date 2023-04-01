package leetcode

import (
	"testing"
)

type question1648 struct {
	para1648
	ans1648
}

// para 是参数
// one 代表第一个参数
type para1648 struct {
	inventory []int
	orders    int
}

// ans 是答案
// one 代表第一个答案
type ans1648 struct {
	one int
}

func Benchmark_Problem1648(b *testing.B) {

	qs := []question1648{

		{
			para1648{[]int{2, 3, 3, 4, 5}, 4},
			ans1648{16},
		},

		{
			para1648{[]int{2, 5}, 4},
			ans1648{14},
		},

		{
			para1648{[]int{3, 5}, 6},
			ans1648{19},
		},

		{
			para1648{[]int{2, 8, 4, 10, 6}, 20},
			ans1648{110},
		},

		{
			para1648{[]int{1000000000}, 1000000000},
			ans1648{21},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1648, q.para1648
				(maxProfit(p.inventory, p.orders))
			}
		}
	}
}

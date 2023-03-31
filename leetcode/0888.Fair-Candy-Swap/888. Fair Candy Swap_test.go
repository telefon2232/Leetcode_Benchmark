package leetcode

import (
	"testing"
)

type question888 struct {
	para888
	ans888
}

// para 是参数
// one 代表第一个参数
type para888 struct {
	one []int
	two []int
}

// ans 是答案
// one 代表第一个答案
type ans888 struct {
	one []int
}

func Benchmark_Problem888(b *testing.B) {

	qs := []question888{

		{
			para888{[]int{}, []int{}},
			ans888{[]int{}},
		},

		{
			para888{[]int{1, 1}, []int{2, 2}},
			ans888{[]int{1, 2}},
		},

		{
			para888{[]int{1, 2}, []int{2, 3}},
			ans888{[]int{1, 2}},
		},

		{
			para888{[]int{2}, []int{1, 3}},
			ans888{[]int{2, 3}},
		},

		{
			para888{[]int{1, 2, 5}, []int{2, 4}},
			ans888{[]int{5, 4}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		_, p := q.ans888, q.para888
		(fairCandySwap(p.one, p.two))
	}
}

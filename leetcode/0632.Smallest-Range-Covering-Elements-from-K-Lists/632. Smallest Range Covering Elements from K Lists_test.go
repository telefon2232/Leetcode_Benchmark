package leetcode

import (
	"testing"
)

type question632 struct {
	para632
	ans632
}

// para 是参数
// one 代表第一个参数
type para632 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans632 struct {
	one []int
}

func Benchmark_Problem632(b *testing.B) {

	qs := []question632{

		{
			para632{[][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}},
			ans632{[]int{20, 24}},
		},
	}

	for _, q := range qs {
		_, p := q.ans632, q.para632
		(smallestRange(p.one))
	}
}

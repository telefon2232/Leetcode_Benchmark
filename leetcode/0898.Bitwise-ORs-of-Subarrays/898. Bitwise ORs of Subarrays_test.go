package leetcode

import (
	"testing"
)

type question898 struct {
	para898
	ans898
}

// para 是参数
// one 代表第一个参数
type para898 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans898 struct {
	one int
}

func Benchmark_Problem898(b *testing.B) {

	qs := []question898{

		{
			para898{[]int{0}},
			ans898{1},
		},

		{
			para898{[]int{1, 1, 2}},
			ans898{3},
		},

		{
			para898{[]int{1, 2, 4}},
			ans898{6},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans898, q.para898
				(subarrayBitwiseORs(p.one))
			}
		}
	}
}

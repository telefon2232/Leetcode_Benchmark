package leetcode

import (
	"testing"
)

type question421 struct {
	para421
	ans421
}

// para 是参数
// one 代表第一个参数
type para421 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans421 struct {
	one int
}

func Benchmark_Problem421(b *testing.B) {

	qs := []question421{

		{
			para421{[]int{3, 10, 5, 25, 2, 8}},
			ans421{28},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans421, q.para421
				(findMaximumXOR(p.one))
			}
		}
	}
}

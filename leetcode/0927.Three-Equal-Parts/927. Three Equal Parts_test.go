package leetcode

import (
	"testing"
)

type question927 struct {
	para927
	ans927
}

// para 是参数
// one 代表第一个参数
type para927 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans927 struct {
	one []int
}

func Benchmark_Problem927(b *testing.B) {

	qs := []question927{

		{
			para927{[]int{1, 0, 1, 0, 1}},
			ans927{[]int{0, 3}},
		},

		{
			para927{[]int{1, 1, 0, 1, 1}},
			ans927{[]int{-1, -1}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans927, q.para927
				(threeEqualParts(p.one))
			}
		}
	}
}

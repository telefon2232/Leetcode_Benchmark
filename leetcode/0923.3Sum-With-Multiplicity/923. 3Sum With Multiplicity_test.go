package leetcode

import (
	"testing"
)

type question923 struct {
	para923
	ans923
}

// para 是参数
// one 代表第一个参数
type para923 struct {
	a []int
	t int
}

// ans 是答案
// one 代表第一个答案
type ans923 struct {
	one int
}

func Benchmark_Problem923(b *testing.B) {

	qs := []question923{

		{
			para923{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8},
			ans923{20},
		},

		{
			para923{[]int{1, 1, 2, 2, 2, 2}, 5},
			ans923{12},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans923, q.para923
		(threeSumMulti(p.a, p.t))
	}
}}}

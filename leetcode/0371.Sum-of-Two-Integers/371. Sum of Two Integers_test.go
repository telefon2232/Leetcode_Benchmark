package leetcode

import (
	"testing"
)

type question371 struct {
	para371
	ans371
}

// para 是参数
// one 代表第一个参数
type para371 struct {
	a int
	b int
}

// ans 是答案
// one 代表第一个答案
type ans371 struct {
	one int
}

func Benchmark_Problem371(b *testing.B) {

	qs := []question371{

		{
			para371{1, 2},
			ans371{3},
		},

		{
			para371{-2, 3},
			ans371{1},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans371, q.para371
		(getSum(p.a, p.b))
	}
}}}

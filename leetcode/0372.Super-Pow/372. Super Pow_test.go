package leetcode

import (
	"testing"
)

type question372 struct {
	para372
	ans372
}

// para 是参数
// one 代表第一个参数
type para372 struct {
	a int
	b []int
}

// ans 是答案
// one 代表第一个答案
type ans372 struct {
	one int
}

func Benchmark_Problem372(b *testing.B) {

	qs := []question372{

		{
			para372{2, []int{3}},
			ans372{8},
		},

		{
			para372{2, []int{1, 0}},
			ans372{1024},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans372, q.para372
		(superPow(p.a, p.b))
	}
}}}

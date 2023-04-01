package leetcode

import (
	"testing"
)

type question509 struct {
	para509
	ans509
}

// para 是参数
// one 代表第一个参数
type para509 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans509 struct {
	one int
}

func Benchmark_Problem509(b *testing.B) {

	qs := []question509{

		{
			para509{1},
			ans509{1},
		},

		{
			para509{2},
			ans509{1},
		},

		{
			para509{3},
			ans509{2},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans509, q.para509
		(fib(p.one))
	}
}}}

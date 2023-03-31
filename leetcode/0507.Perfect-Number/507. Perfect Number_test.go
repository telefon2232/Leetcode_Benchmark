package leetcode

import (
	"testing"
)

type question507 struct {
	para507
	ans507
}

// para 是参数
// one 代表第一个参数
type para507 struct {
	num int
}

// ans 是答案
// one 代表第一个答案
type ans507 struct {
	one bool
}

func Benchmark_Problem507(b *testing.B) {

	qs := []question507{

		{
			para507{28},
			ans507{true},
		},

		{
			para507{496},
			ans507{true},
		},

		{
			para507{500},
			ans507{false},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		_, p := q.ans507, q.para507
		(checkPerfectNumber(p.num))
	}
}

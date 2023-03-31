package leetcode

import (
	"testing"
)

type question326 struct {
	para326
	ans326
}

// para 是参数
// one 代表第一个参数
type para326 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans326 struct {
	one bool
}

func Benchmark_Problem326(b *testing.B) {

	qs := []question326{

		{
			para326{27},
			ans326{true},
		},

		{
			para326{0},
			ans326{false},
		},

		{
			para326{9},
			ans326{true},
		},

		{
			para326{45},
			ans326{false},
		},
	}

	for _, q := range qs {
		_, p := q.ans326, q.para326
		(isPowerOfThree(p.one))
	}
}

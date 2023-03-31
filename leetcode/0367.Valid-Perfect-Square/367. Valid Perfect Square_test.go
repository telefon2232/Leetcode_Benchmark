package leetcode

import (
	"testing"
)

type question367 struct {
	para367
	ans367
}

// para 是参数
// one 代表第一个参数
type para367 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans367 struct {
	one bool
}

func Benchmark_Problem367(b *testing.B) {

	qs := []question367{

		{
			para367{1},
			ans367{true},
		},

		{
			para367{2},
			ans367{false},
		},

		{
			para367{3},
			ans367{false},
		},

		{
			para367{4},
			ans367{true},
		},

		{
			para367{5},
			ans367{false},
		},

		{
			para367{6},
			ans367{false},
		},

		{
			para367{104976},
			ans367{true},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		_, p := q.ans367, q.para367
		(isPerfectSquare(p.one))
	}
}

package leetcode

import (
	"testing"
)

type question1758 struct {
	para1758
	ans1758
}

// para 是参数
// one 代表第一个参数
type para1758 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans1758 struct {
	one int
}

func Benchmark_Problem1758(b *testing.B) {

	qs := []question1758{

		{
			para1758{"0100"},
			ans1758{1},
		},

		{
			para1758{"10"},
			ans1758{0},
		},

		{
			para1758{"1111"},
			ans1758{2},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans1758, q.para1758
		(minOperations(p.s))
	}
}}}

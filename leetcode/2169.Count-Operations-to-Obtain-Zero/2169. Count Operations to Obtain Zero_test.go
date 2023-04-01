package leetcode

import (
	"testing"
)

type question2169 struct {
	para2169
	ans2169
}

// para 是参数
// one 代表第一个参数
type para2169 struct {
	num1 int
	num2 int
}

// ans 是答案
// one 代表第一个答案
type ans2169 struct {
	one int
}

func Benchmark_Problem2169(b *testing.B) {

	qs := []question2169{

		{
			para2169{2, 3},
			ans2169{3},
		},

		{
			para2169{10, 10},
			ans2169{1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans2169, q.para2169
		(countOperations(p.num1, p.num2))
	}
}}}

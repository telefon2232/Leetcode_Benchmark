package leetcode

import (
	"testing"
)

type question791 struct {
	para791
	ans791
}

// para 是参数
// one 代表第一个参数
type para791 struct {
	order string
	str   string
}

// ans 是答案
// one 代表第一个答案
type ans791 struct {
	one string
}

func Benchmark_Problem791(b *testing.B) {

	qs := []question791{

		{
			para791{"cba", "abcd"},
			ans791{"cbad"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans791, q.para791
		(customSortString(p.order, p.str))
	}
}}}

package leetcode

import (
	"testing"
)

type question22 struct {
	para22
	ans22
}

// para 是参数
// one 代表第一个参数
type para22 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans22 struct {
	one []string
}

func Benchmark_Problem22(b *testing.B) {

	qs := []question22{

		{
			para22{3},
			ans22{[]string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans22, q.para22
		(generateParenthesis(p.one))
	}
}}}

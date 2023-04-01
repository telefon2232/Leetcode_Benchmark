package leetcode

import (
	"testing"
)

type question1002 struct {
	para1002
	ans1002
}

// para 是参数
// one 代表第一个参数
type para1002 struct {
	one []string
}

// ans 是答案
// one 代表第一个答案
type ans1002 struct {
	one []string
}

func Benchmark_Problem1002(b *testing.B) {

	qs := []question1002{

		{
			para1002{[]string{"bella", "label", "roller"}},
			ans1002{[]string{"e", "l", "l"}},
		},

		{
			para1002{[]string{"cool", "lock", "cook"}},
			ans1002{[]string{"c", "o"}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans1002, q.para1002
		(commonChars(p.one))
	}
}}}

package leetcode

import (
	"testing"
)

type question784 struct {
	para784
	ans784
}

// para 是参数
// one 代表第一个参数
type para784 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans784 struct {
	one []string
}

func Benchmark_Problem784(b *testing.B) {

	qs := []question784{

		{
			para784{"mQe"},
			ans784{[]string{"mqe", "mqE", "mQe", "mQE", "Mqe", "MqE", "MQe", "MQE"}},
		},

		{
			para784{"C"},
			ans784{[]string{"c", "C"}},
		},

		{
			para784{"a1b2"},
			ans784{[]string{"a1b2", "a1B2", "A1b2", "A1B2"}},
		},

		{
			para784{"3z4"},
			ans784{[]string{"3z4", "3Z4"}},
		},

		{
			para784{"12345"},
			ans784{[]string{"12345"}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans784, q.para784
		(letterCasePermutation1(p.one))
	}
}}}

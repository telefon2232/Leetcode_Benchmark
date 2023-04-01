package leetcode

import (
	"testing"
)

type question880 struct {
	para880
	ans880
}

// para 是参数
// one 代表第一个参数
type para880 struct {
	s string
	k int
}

// ans 是答案
// one 代表第一个答案
type ans880 struct {
	one string
}

func Benchmark_Problem880(b *testing.B) {

	qs := []question880{

		{
			para880{"aw4eguc6cs", 41},
			ans880{"a"},
		},

		{
			para880{"vk6u5xhq9v", 554},
			ans880{"h"},
		},

		{
			para880{"leet2code3", 10},
			ans880{"o"},
		},

		{
			para880{"ha22", 5},
			ans880{"h"},
		},

		{
			para880{"a2345678999999999999999", 1},
			ans880{"a"},
		},

		{
			para880{"y959q969u3hb22odq595", 222280369},
			ans880{"q"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans880, q.para880
		(decodeAtIndex(p.s, p.k))
	}
}}}

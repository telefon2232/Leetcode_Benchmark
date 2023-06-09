package leetcode

import (
	"testing"
)

type question394 struct {
	para394
	ans394
}

// para 是参数
// one 代表第一个参数
type para394 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans394 struct {
	one string
}

func Benchmark_Problem394(b *testing.B) {

	qs := []question394{

		{
			para394{"10[a]"},
			ans394{"aaaaaaaaaa"},
		},

		{
			para394{"3[a]2[bc]"},
			ans394{"aaabcbc"},
		},

		{
			para394{"3[a2[c]]"},
			ans394{"accaccacc"},
		},

		{
			para394{"2[abc]3[cd]ef"},
			ans394{"abcabccdcdcdef"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans394, q.para394
				(decodeString(p.s))
			}
		}
	}
}

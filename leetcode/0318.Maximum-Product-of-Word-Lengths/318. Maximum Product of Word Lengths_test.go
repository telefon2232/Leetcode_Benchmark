package leetcode

import (
	"testing"
)

type question318 struct {
	para318
	ans318
}

// para 是参数
// one 代表第一个参数
type para318 struct {
	one []string
}

// ans 是答案
// one 代表第一个答案
type ans318 struct {
	one int
}

func Benchmark_Problem318(b *testing.B) {

	qs := []question318{

		{
			para318{[]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}},
			ans318{16},
		},

		{
			para318{[]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}},
			ans318{4},
		},

		{
			para318{[]string{"a", "aa", "aaa", "aaaa"}},
			ans318{0},
		},
	}

	for _, q := range qs {
		_, p := q.ans318, q.para318
		(maxProduct318(p.one))
	}
}

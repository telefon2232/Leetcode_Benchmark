package leetcode

import (
	"testing"
)

type question242 struct {
	para242
	ans242
}

// para 是参数
// one 代表第一个参数
type para242 struct {
	one string
	two string
}

// ans 是答案
// one 代表第一个答案
type ans242 struct {
	one bool
}

func Benchmark_Problem242(b *testing.B) {

	qs := []question242{

		{
			para242{"", ""},
			ans242{true},
		},
		{
			para242{"", "1"},
			ans242{false},
		},

		{
			para242{"anagram", "nagaram"},
			ans242{true},
		},

		{
			para242{"rat", "car"},
			ans242{false},
		},

		{
			para242{"a", "ab"},
			ans242{false},
		},

		{
			para242{"ab", "a"},
			ans242{false},
		},

		{
			para242{"aa", "bb"},
			ans242{false},
		},
	}

	for _, q := range qs {
		_, p := q.ans242, q.para242
		(isAnagram(p.one, p.two))
	}
}

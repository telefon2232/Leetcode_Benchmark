package leetcode

import (
	"testing"
)

type question1160 struct {
	para1160
	ans1160
}

// para 是参数
// one 代表第一个参数
type para1160 struct {
	words []string
	chars string
}

// ans 是答案
// one 代表第一个答案
type ans1160 struct {
	one int
}

func Benchmark_Problem1160(b *testing.B) {

	qs := []question1160{

		{
			para1160{[]string{"cat", "bt", "hat", "tree"}, "atach"},
			ans1160{6},
		},

		{
			para1160{[]string{"hello", "world", "leetcode"}, "welldonehoneyr"},
			ans1160{10},
		},
	}

	for _, q := range qs {
		_, p := q.ans1160, q.para1160
		(countCharacters(p.words, p.chars))
	}
}

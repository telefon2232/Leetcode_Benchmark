package leetcode

import (
	"testing"
)

type question828 struct {
	para828
	ans828
}

// para 是参数
// one 代表第一个参数
type para828 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans828 struct {
	one int
}

func Benchmark_Problem828(b *testing.B) {

	qs := []question828{

		{
			para828{"BABABBABAA"},
			ans828{35},
		},

		{
			para828{"ABC"},
			ans828{10},
		},

		{
			para828{"ABA"},
			ans828{8},
		},

		{
			para828{"ABAB"},
			ans828{12},
		},
	}

	for _, q := range qs {
		_, p := q.ans828, q.para828
		(uniqueLetterString(p.one))
	}
}

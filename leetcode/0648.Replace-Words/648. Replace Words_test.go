package leetcode

import (
	"testing"
)

type question648 struct {
	para648
	ans648
}

// para 是参数
// one 代表第一个参数
type para648 struct {
	one []string
	s   string
}

// ans 是答案
// one 代表第一个答案
type ans648 struct {
	one string
}

func Benchmark_Problem648(b *testing.B) {

	qs := []question648{

		{
			para648{[]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"},
			ans648{"the cat was rat by the bat"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans648, q.para648
				(replaceWords(p.one, p.s))
			}
		}
	}
}

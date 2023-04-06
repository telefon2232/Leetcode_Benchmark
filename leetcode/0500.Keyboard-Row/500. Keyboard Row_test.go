package leetcode

import (
	"testing"
)

type question500 struct {
	para500
	ans500
}

// para 是参数
// one 代表第一个参数
type para500 struct {
	one []string
}

// ans 是答案
// one 代表第一个答案
type ans500 struct {
	one []string
}

func Benchmark_Problem500(b *testing.B) {

	qs := []question500{

		{
			para500{[]string{"Hello", "Alaska", "Dad", "Peace"}},
			ans500{[]string{"Alaska", "Dad"}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans500, q.para500
				(findWords500(p.one))
			}
		}
	}
}

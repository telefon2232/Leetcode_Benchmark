package leetcode

import (
	"testing"
)

type question151 struct {
	para151
	ans151
}

// para 是参数
// one 代表第一个参数
type para151 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans151 struct {
	one string
}

func Benchmark_Problem151(b *testing.B) {

	qs := []question151{

		{
			para151{"the sky is blue"},
			ans151{"blue is sky the"},
		},

		{
			para151{"  hello world!  "},
			ans151{"world! hello"},
		},
		{
			para151{"a good   example"},
			ans151{"example good a"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans151, q.para151
				(reverseWords151(p.one))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question820 struct {
	para820
	ans820
}

// para 是参数
// one 代表第一个参数
type para820 struct {
	words []string
}

// ans 是答案
// one 代表第一个答案
type ans820 struct {
	one int
}

func Benchmark_Problem820(b *testing.B) {

	qs := []question820{

		{
			para820{[]string{"time", "me", "bell"}},
			ans820{10},
		},

		{
			para820{[]string{"t"}},
			ans820{2},
		},
	}

	for _, q := range qs {
		_, p := q.ans820, q.para820
		(minimumLengthEncoding(p.words))
	}
}

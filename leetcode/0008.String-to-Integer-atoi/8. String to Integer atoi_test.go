package leetcode

import (
	"testing"
)

type question8 struct {
	para8
	ans8
}

// para 是参数
// one 代表第一个参数
type para8 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans8 struct {
	one int
}

func Benchmark_Problem8(b *testing.B) {

	qs := []question8{

		{
			para8{"42"},
			ans8{42},
		},

		{
			para8{"   -42"},
			ans8{-42},
		},

		{
			para8{"4193 with words"},
			ans8{4193},
		},

		{
			para8{"words and 987"},
			ans8{0},
		},

		{
			para8{"-91283472332"},
			ans8{-2147483648},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans8, q.para8
				(myAtoi(p.one))
			}
		}
	}
}

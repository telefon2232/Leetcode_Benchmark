package leetcode

import (
	"testing"
)

type question405 struct {
	para405
	ans405
}

// para 是参数
// one 代表第一个参数
type para405 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans405 struct {
	one string
}

func Benchmark_Problem405(b *testing.B) {

	qs := []question405{

		{
			para405{26},
			ans405{"1a"},
		},

		{
			para405{-1},
			ans405{"ffffffff"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans405, q.para405
				(toHex(p.one))
			}
		}
	}
}

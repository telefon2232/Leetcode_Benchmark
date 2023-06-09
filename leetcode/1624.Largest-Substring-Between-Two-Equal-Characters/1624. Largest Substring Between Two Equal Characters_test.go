package leetcode

import (
	"testing"
)

type question1624 struct {
	para1624
	ans1624
}

// para 是参数
// one 代表第一个参数
type para1624 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans1624 struct {
	one int
}

func Benchmark_Problem1624(b *testing.B) {

	qs := []question1624{

		{
			para1624{"aa"},
			ans1624{0},
		},

		{
			para1624{"abca"},
			ans1624{2},
		},

		{
			para1624{"cbzxy"},
			ans1624{-1},
		},

		{
			para1624{"cabbac"},
			ans1624{4},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1624, q.para1624
				(maxLengthBetweenEqualCharacters(p.s))
			}
		}
	}
}

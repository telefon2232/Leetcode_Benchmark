package leetcode

import (
	"testing"
)

type question319 struct {
	para319
	ans319
}

// para 是参数
type para319 struct {
	n int
}

// ans 是答案
type ans319 struct {
	ans int
}

func Benchmark_Problem319(b *testing.B) {

	qs := []question319{

		{
			para319{3},
			ans319{1},
		},

		{
			para319{0},
			ans319{0},
		},

		{
			para319{1},
			ans319{1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans319, q.para319
				(bulbSwitch(p.n))
			}
		}
	}
}

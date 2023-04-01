package leetcode

import (
	"testing"
)

type question6 struct {
	para6
	ans6
}

// para 是参数
// one 代表第一个参数
type para6 struct {
	s       string
	numRows int
}

// ans 是答案
// one 代表第一个答案
type ans6 struct {
	one string
}

func Benchmark_Problem6(b *testing.B) {

	qs := []question6{

		{
			para6{"PAYPALISHIRING", 3},
			ans6{"PAHNAPLSIIGYIR"},
		},

		{
			para6{"PAYPALISHIRING", 4},
			ans6{"PINALSIGYAHRPI"},
		},

		{
			para6{"A", 1},
			ans6{"A"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans6, q.para6
		(convert(p.s, p.numRows))
	}
}}}

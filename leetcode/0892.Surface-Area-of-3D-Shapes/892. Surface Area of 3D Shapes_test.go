package leetcode

import (
	"testing"
)

type question892 struct {
	para892
	ans892
}

// para 是参数
// one 代表第一个参数
type para892 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans892 struct {
	one int
}

func Benchmark_Problem892(b *testing.B) {

	qs := []question892{

		{
			para892{[][]int{{2}}},
			ans892{10},
		},

		{
			para892{[][]int{{1, 2}, {3, 4}}},
			ans892{34},
		},

		{
			para892{[][]int{{1, 0}, {0, 2}}},
			ans892{16},
		},

		{
			para892{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}},
			ans892{32},
		},

		{
			para892{[][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}},
			ans892{46},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans892, q.para892
		(surfaceArea(p.one))
	}
}}}

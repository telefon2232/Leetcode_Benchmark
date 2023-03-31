package leetcode

import (
	"testing"
)

type question200 struct {
	para200
	ans200
}

// para 是参数
// one 代表第一个参数
type para200 struct {
	one [][]byte
}

// ans 是答案
// one 代表第一个答案
type ans200 struct {
	one int
}

func Benchmark_Problem200(b *testing.B) {

	qs := []question200{

		{
			para200{[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			}},
			ans200{1},
		},

		{
			para200{[][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			}},
			ans200{3},
		},
	}

	for _, q := range qs {
		_, p := q.ans200, q.para200
		(numIslands(p.one))
	}
}

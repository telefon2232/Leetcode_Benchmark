package leetcode

import (
	"fmt"
	"testing"
)

type question999 struct {
	para999
	ans999
}

// para 是参数
// one 代表第一个参数
type para999 struct {
	one [][]byte
}

// ans 是答案
// one 代表第一个答案
type ans999 struct {
	one int
}

func Benchmark_Problem999(b *testing.B) {

	qs := []question999{

		{
			para999{[][]byte{
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
				{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
				{'.', 'p', 'B', 'R', 'B', 'p', '.', '.'},
				{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
				{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
			}},
			ans999{0},
		},

		{
			para999{[][]byte{
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', 'p', '.', '.', '.', '.'},
				{'.', '.', '.', 'R', '.', '.', '.', 'p'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', 'p', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
			}},
			ans999{3},
		},

		{
			para999{[][]byte{
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', 'p', '.', '.', '.', '.'},
				{'.', '.', '.', 'p', '.', '.', '.', '.'},
				{'p', 'p', '.', 'R', '.', 'p', 'B', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', 'B', '.', '.', '.', '.'},
				{'.', '.', '.', 'p', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
			}},
			ans999{3},
		},
	}


	for _, q := range qs {
		_, p := q.ans999, q.para999
		fmt.Printf("【input】:%v       【output】:%v\n", p, numRookCaptures(p.one))
	}
}

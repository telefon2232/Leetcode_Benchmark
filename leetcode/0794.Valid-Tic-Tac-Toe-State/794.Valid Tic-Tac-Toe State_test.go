package leetcode

import (
	"fmt"
	"testing"
)

type question794 struct {
	para794
	ans794
}

// para 是参数
type para794 struct {
	board []string
}

// ans 是答案
type ans794 struct {
	ans bool
}

func Benchmark_Problem794(b *testing.B) {

	qs := []question794{

		{
			para794{[]string{"O  ", "   ", "   "}},
			ans794{false},
		},

		{
			para794{[]string{"XOX", " X ", "   "}},
			ans794{false},
		},

		{
			para794{[]string{"XXX", "   ", "OOO"}},
			ans794{false},
		},

		{
			para794{[]string{"XOX", "O O", "XOX"}},
			ans794{true},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans794, q.para794
				fmt.Printf("【input】:%v    【output】:%v\n", p.board, validTicTacToe(p.board))
			}
		}
	}
}

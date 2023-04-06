package leetcode

import (
	"fmt"
	"testing"
)

type question2038 struct {
	para2038
	ans2038
}

// para 是参数
type para2038 struct {
	colors string
}

// ans 是答案
type ans2038 struct {
	ans bool
}

func Benchmark_Problem2038(b *testing.B) {

	qs := []question2038{

		{
			para2038{"AAABABB"},
			ans2038{true},
		},

		{
			para2038{"AA"},
			ans2038{false},
		},

		{
			para2038{"ABBBBBBBAAA"},
			ans2038{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans2038, q.para2038
				fmt.Printf("【input】:%v      ", p.colors)
				(winnerOfGame(p.colors))
			}
		}
	}
}

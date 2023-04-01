package leetcode

import (
	"testing"
)

type question909 struct {
	para909
	ans909
}

// para 是参数
// one 代表第一个参数
type para909 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans909 struct {
	one int
}

func Benchmark_Problem909(b *testing.B) {
	qs := []question909{
		{
			para909{[][]int{
				{-1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, 35, -1, -1, 13, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, 15, -1, -1, -1, -1},
			}},
			ans909{4},
		},
	}
	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans909, q.para909
		(snakesAndLadders(p.one))
	}
}}}

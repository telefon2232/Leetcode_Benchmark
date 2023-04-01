package leetcode

import (
	"testing"
)

type question461 struct {
	para461
	ans461
}

// para 是参数
// one 代表第一个参数
type para461 struct {
	x int
	y int
}

// ans 是答案
// one 代表第一个答案
type ans461 struct {
	one int
}

func Benchmark_Problem461(b *testing.B) {

	qs := []question461{

		{
			para461{1, 4},
			ans461{2},
		},

		{
			para461{1, 1},
			ans461{0},
		},

		{
			para461{1, 3},
			ans461{1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans461, q.para461
		(hammingDistance(p.x, p.y))
	}
}}}

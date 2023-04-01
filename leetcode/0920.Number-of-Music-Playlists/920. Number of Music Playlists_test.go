package leetcode

import (
	"testing"
)

type question920 struct {
	para920
	ans920
}

// para 是参数
// one 代表第一个参数
type para920 struct {
	N int
	L int
	K int
}

// ans 是答案
// one 代表第一个答案
type ans920 struct {
	one int
}

func Benchmark_Problem920(b *testing.B) {

	qs := []question920{

		{
			para920{3, 3, 1},
			ans920{6},
		},

		{
			para920{2, 3, 0},
			ans920{6},
		},

		{
			para920{2, 3, 1},
			ans920{2},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans920, q.para920
		(numMusicPlaylists(p.N, p.L, p.K))
	}
}}}

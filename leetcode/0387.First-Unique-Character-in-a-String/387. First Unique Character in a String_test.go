package leetcode

import (
	"testing"
)

type question387 struct {
	para387
	ans387
}

// para 是参数
// one 代表第一个参数
type para387 struct {
	n string
}

// ans 是答案
// one 代表第一个答案
type ans387 struct {
	one int
}

func Benchmark_Problem387(b *testing.B) {

	qs := []question387{

		{
			para387{"leetcode"},
			ans387{0},
		},

		{
			para387{"loveleetcode"},
			ans387{2},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans387, q.para387
		(firstUniqChar(p.n))
	}
}}}

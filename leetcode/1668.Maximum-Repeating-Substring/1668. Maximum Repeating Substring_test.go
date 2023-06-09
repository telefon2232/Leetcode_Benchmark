package leetcode

import (
	"testing"
)

type question1668 struct {
	para1668
	ans1668
}

// para 是参数
// one 代表第一个参数
type para1668 struct {
	sequence string
	word     string
}

// ans 是答案
// one 代表第一个答案
type ans1668 struct {
	one int
}

func Benchmark_Problem1668(b *testing.B) {

	qs := []question1668{

		{
			para1668{"ababc", "ab"},
			ans1668{2},
		},

		{
			para1668{"ababc", "ba"},
			ans1668{1},
		},

		{
			para1668{"ababc", "ac"},
			ans1668{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1668, q.para1668
				(maxRepeating(p.sequence, p.word))
			}
		}
	}
}

package leetcode

import (
	"testing"
)

type question763 struct {
	para763
	ans763
}

// para 是参数
// one 代表第一个参数
type para763 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans763 struct {
	one []int
}

func Benchmark_Problem763(b *testing.B) {

	qs := []question763{

		{
			para763{"ababcbacadefegdehijhklij"},
			ans763{[]int{9, 7, 8}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans763, q.para763
				(partitionLabels(p.one))
			}
		}
	}
}

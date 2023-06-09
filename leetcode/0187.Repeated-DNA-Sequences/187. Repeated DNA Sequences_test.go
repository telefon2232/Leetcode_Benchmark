package leetcode

import (
	"testing"
)

type question187 struct {
	para187
	ans187
}

// para 是参数
// one 代表第一个参数
type para187 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans187 struct {
	one []string
}

func Benchmark_Problem187(b *testing.B) {

	qs := []question187{

		{
			para187{"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"},
			ans187{[]string{"AAAAACCCCC", "CCCCCAAAAA"}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans187, q.para187
				(findRepeatedDnaSequences(p.one))
			}
		}
	}
}

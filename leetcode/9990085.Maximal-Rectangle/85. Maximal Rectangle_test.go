package leetcode

import (
	"testing"
)

type question85 struct {
	para85
	ans85
}

// para 是参数
// one 代表第一个参数
type para85 struct {
	one [][]byte
}

// ans 是答案
// one 代表第一个答案
type ans85 struct {
	one int
}

func Benchmark_Problem85(b *testing.B) {

	qs := []question85{

		{
			para85{[][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}},
			ans85{6},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans85, q.para85
				(maximalRectangle(p.one))
			}
		}
	}
}

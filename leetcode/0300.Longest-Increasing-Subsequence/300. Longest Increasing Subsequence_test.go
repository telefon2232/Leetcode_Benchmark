package leetcode

import (
	"testing"
)

type question300 struct {
	para300
	ans300
}

// para 是参数
// one 代表第一个参数
type para300 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans300 struct {
	one int
}

func Benchmark_Problem300(b *testing.B) {

	qs := []question300{

		{
			para300{[]int{10, 9, 2, 5, 3, 7, 101, 18}},
			ans300{4},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans300, q.para300
		(lengthOfLIS(p.one))
	}
}}}

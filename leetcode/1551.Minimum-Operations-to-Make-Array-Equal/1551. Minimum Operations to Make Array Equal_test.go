package leetcode

import (
	"fmt"
	"testing"
)

type question1551 struct {
	para1551
	ans1551
}

// para 是参数
// one 代表第一个参数
type para1551 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans1551 struct {
	one int
}

func Benchmark_Problem1551(b *testing.B) {

	qs := []question1551{

		{
			para1551{3},
			ans1551{2},
		},

		{
			para1551{6},
			ans1551{9},
		},

		{
			para1551{534},
			ans1551{71289},
		},
	}


	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans1551, q.para1551
		fmt.Printf("【input】:%v      【output】:%v      \n", p, minOperations(p.n))
	}
}}}

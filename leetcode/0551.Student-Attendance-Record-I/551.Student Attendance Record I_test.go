package leetcode

import (
	"fmt"
	"testing"
)

type question551 struct {
	para551
	ans551
}

// para 是参数
type para551 struct {
	s string
}

// ans 是答案
type ans551 struct {
	ans bool
}

func Benchmark_Problem551(b *testing.B) {

	qs := []question551{

		{
			para551{"PPALLP"},
			ans551{true},
		},

		{
			para551{"PPALLL"},
			ans551{false},
		},
	}


	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans551, q.para551
		fmt.Printf("【input】:%v      【output】:%v      \n", p, checkRecord(p.s))
	}
}}}

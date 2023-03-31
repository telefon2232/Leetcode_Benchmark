package leetcode

import (
	"fmt"
	"testing"
)

type question91 struct {
	para91
	ans91
}

// para 是参数
// one 代表第一个参数
type para91 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans91 struct {
	one int
}

func Benchmark_Problem91(b *testing.B) {

	qs := []question91{

		{
			para91{"12"},
			ans91{2},
		},

		{
			para91{"226"},
			ans91{3},
		},
	}


	for _, q := range qs {
		_, p := q.ans91, q.para91
		fmt.Printf("【input】:%v       【output】:%v\n", p, numDecodings(p.one))
	}
}

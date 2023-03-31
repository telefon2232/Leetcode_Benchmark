package leetcode

import (
	"fmt"
	"testing"
)

type question204 struct {
	para204
	ans204
}

// para 是参数
// one 代表第一个参数
type para204 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans204 struct {
	one int
}

func Benchmark_Problem204(b *testing.B) {

	qs := []question204{

		{
			para204{10},
			ans204{4},
		},

		{
			para204{100},
			ans204{25},
		},

		{
			para204{1000},
			ans204{168},
		},
	}


	for _, q := range qs {
		_, p := q.ans204, q.para204
		fmt.Printf("【input】:%v       【output】:%v\n", p, countPrimes(p.one))
	}
}

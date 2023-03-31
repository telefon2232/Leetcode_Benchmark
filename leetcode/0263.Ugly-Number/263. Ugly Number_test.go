package leetcode

import (
	"fmt"
	"testing"
)

type question263 struct {
	para263
	ans263
}

// para 是参数
// one 代表第一个参数
type para263 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans263 struct {
	one bool
}

func Benchmark_Problem263(b *testing.B) {

	qs := []question263{

		{
			para263{6},
			ans263{true},
		},

		{
			para263{8},
			ans263{true},
		},

		{
			para263{14},
			ans263{false},
		},
	}


	for _, q := range qs {
		_, p := q.ans263, q.para263
		fmt.Printf("【input】:%v       【output】:%v\n", p, isUgly(p.one))
	}
}

package leetcode

import (
	"fmt"
	"testing"
)

type question470 struct {
	para470
	ans470
}

// para 是参数
// one 代表第一个参数
type para470 struct {
}

// ans 是答案
// one 代表第一个答案
type ans470 struct {
	one int
}

func Benchmark_Problem470(b *testing.B) {

	qs := []question470{

		{
			para470{},
			ans470{2},
		},

		{
			para470{},
			ans470{0},
		},

		{
			para470{},
			ans470{1},
		},
	}


	for _, q := range qs {
		_, p := q.ans470, q.para470
		fmt.Printf("【input】:%v       【output】:%v\n", p, rand10())
	}
}

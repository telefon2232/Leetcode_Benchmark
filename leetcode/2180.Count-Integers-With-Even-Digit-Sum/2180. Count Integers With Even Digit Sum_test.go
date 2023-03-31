package leetcode

import (
	"fmt"
	"testing"
)

type question2180 struct {
	para2180
	ans2180
}

// para 是参数
// one 代表第一个参数
type para2180 struct {
	target int
}

// ans 是答案
// one 代表第一个答案
type ans2180 struct {
	one int
}

func Benchmark_Problem1(b *testing.B) {

	qs := []question2180{
		{
			para2180{4},
			ans2180{2},
		},

		{
			para2180{30},
			ans2180{14},
		},
	}


	for _, q := range qs {
		_, p := q.ans2180, q.para2180
		fmt.Printf("【input】:%v       【output】:%v\n", p, countEven(p.target))
	}
}
